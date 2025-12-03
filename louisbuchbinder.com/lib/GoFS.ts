interface IGoDirEntry {
  name(): string;
  isDir(): boolean;
  type(): null;
  info(): IGoFileInfo | Error;
}

interface IGoReadDirFile extends IGoFile {
  readDir(n: number): IGoDirEntry[];
}

interface IGoFS {
  open(p: string): Promise<IGoFile>;
}

class GoDirEntry implements IGoDirEntry {
  _info: IGoFileInfo;

  constructor(info: IGoFileInfo) {
    this._info = info;
  }

  name(): string {
    return this._info.name();
  }

  isDir(): boolean {
    return this._info.isDir();
  }

  type(): null {
    return this._info.mode();
  }

  info(): IGoFileInfo | Error {
    return this._info;
  }
}

class GoReadDirFile extends GoFile implements IGoReadDirFile {
  info: GoDirFileInfo;
  begin: number;
  entries: IGoDirEntry[];

  constructor(dirName: string, entries: IGoDirEntry[]) {
    super(null);
    this.begin = 0;
    this.entries = entries;
    this.info = new GoDirFileInfo(dirName);
  }

  async stat(): Promise<IGoFileInfo> {
    return this.info;
  }

  async read(): Promise<Uint8Array<ArrayBufferLike>> {
    return null;
  }

  readDir(n: number): IGoDirEntry[] {
    if (n === -1) {
      this.begin = 0;
      return this.entries;
    }
    if (this.begin >= this.entries.length) {
      return null;
    }
    const entries = this.entries.slice(this.begin, this.begin + n);
    this.begin += n;
    return entries;
  }
}

class GoDirFileInfo implements IGoFileInfo {
  dirName: string;

  constructor(dirName: string) {
    this.dirName = dirName;
  }

  name(): string {
    return this.dirName;
  }

  mode(): null {
    return null;
  }

  size(): number {
    return 0;
  }

  modTime(): number {
    return 0;
  }

  isDir(): boolean {
    return true;
  }

  sys(): null {
    return null;
  }
}

class GoFS implements IGoFS {
  static ErrInvalid = new Error("invalid");
  static ErrNotFound = new Error("not found");

  trie: FSTrie;

  constructor(l: FileList) {
    this.trie = new FSTrie("");
    for (let i = 0; i < l.length; i++) {
      const f = l.item(i);
      if (typeof f.webkitRelativePath === "undefined") {
        throw new Error("ERROR: webkitRelativePath not found for file");
      }
      if (typeof f.webkitRelativePath !== "string") {
        throw new Error(
          "ERROR: webkitRelativePath type violation, expected string, but instead got: " +
            typeof f.webkitRelativePath
        );
      }
      const relPath =
        f.webkitRelativePath.length > 0 ? f.webkitRelativePath : f.name;
      const goFile = new GoFile(f);
      this.trie.add(relPath, goFile);
    }
  }

  async open(p: string): Promise<IGoFile> {
    return this.trie.get(p);
  }
}

class FSTrie {
  prefix: string;
  files: Map<string, GoFile>;
  dirs: Map<string, FSTrie>;

  constructor(p: string) {
    this.prefix = p;
    this.files = new Map();
    this.dirs = new Map();
  }

  add(p: string, file: GoFile) {
    if (p.length === 0) {
      return;
    }
    const parts = p.split("/");
    let current: FSTrie = this;

    for (let i = 0; i < parts.length - 1; i++) {
      if (!current.dirs.has(parts[i])) {
        current.dirs.set(
          parts[i],
          new FSTrie(
            [current.prefix, parts[i]]
              .slice(current.prefix.length === 0 ? 1 : 0)
              .join("/")
          )
        );
      }
      current = current.dirs.get(parts[i]);
    }
    current.files.set(parts[parts.length - 1], file);
  }

  async get(p: string): Promise<IGoFile> {
    if (p.length === 0) {
      throw GoFS.ErrInvalid;
    }
    if (p[0] === "/") {
      // Only relative paths are supported
      throw GoFS.ErrInvalid;
    }
    if (p === ".") {
      return new GoReadDirFile(".", await this.entries());
    }
    const parts = p.split("/");
    let current: FSTrie = this;
    for (let i = 0; i < parts.length - 1; i++) {
      if (!current.dirs.has(parts[i])) {
        throw GoFS.ErrNotFound;
      }
      current = current.dirs.get(parts[i]);
    }
    if (current.files.get(parts[parts.length - 1])) {
      return current.files.get(parts[parts.length - 1]);
    }
    if (current.dirs.get(parts[parts.length - 1])) {
      return new GoReadDirFile(
        parts[parts.length - 1],
        await current.dirs.get(parts[parts.length - 1]).entries()
      );
    }
    throw GoFS.ErrNotFound;
  }

  async entries(): Promise<IGoDirEntry[]> {
    const files = Array.from(this.files.values()).map(
      async (f) => new GoDirEntry(await f.stat())
    );
    const dirs = Array.from(this.dirs.entries()).map(
      async ([k, v]) =>
        new GoDirEntry(await new GoReadDirFile(k, await v.entries()).stat())
    );
    return await Promise.all(Array.prototype.concat(files, dirs));
  }
}
