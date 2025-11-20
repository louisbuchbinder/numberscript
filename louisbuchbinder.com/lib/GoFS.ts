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
  open(p: string): IGoFile | Error;
}

class GoDirEntry implements IGoDirEntry {
  f: IGoFile;

  constructor(f: IGoFile) {
    this.f = f;
  }

  name(): string {
    return this.f.stat().name();
  }

  isDir(): boolean {
    return this.f.stat().isDir();
  }

  type(): null {
    return this.f.stat().mode();
  }

  info(): IGoFileInfo | Error {
    return this.f.stat();
  }
}

class GoReadDirFile implements IGoReadDirFile {
  info: GoDirFileInfo;
  begin: number;
  entries: IGoDirEntry[];

  constructor(dirName: string, entries: IGoDirEntry[]) {
    this.begin = 0;
    this.entries = entries;
    this.info = new GoDirFileInfo(dirName);
  }

  stat(): GoDirFileInfo {
    return this.info;
  }

  async read(): Promise<Uint8Array<ArrayBufferLike>> {
    return null;
  }

  readDir(n: number): IGoDirEntry[] {
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

  modTime(): null {
    return null;
  }

  isDir(): boolean {
    return true;
  }

  sys(): null {
    return null;
  }
}

class GoFS {
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

  open(p: string): IGoFile | Error {
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

  get(p: string): IGoFile | Error {
    if (p.length === 0) {
      return GoFS.ErrInvalid;
    }
    const parts = p.split("/");
    let current: FSTrie = this;
    for (let i = 0; i < parts.length - 1; i++) {
      if (!current.dirs.has(parts[i])) {
        return GoFS.ErrNotFound;
      }
      current = current.dirs.get(parts[i]);
    }
    if (current.files.get(parts[parts.length - 1])) {
      return current.files.get(parts[parts.length - 1]);
    }
    if (current.dirs.get(parts[parts.length - 1])) {
      return new GoReadDirFile(
        parts[parts.length - 1],
        current.dirs.get(parts[parts.length - 1]).entries()
      );
    }
    return GoFS.ErrNotFound;
  }

  entries(): IGoDirEntry[] {
    const files = Array.from(this.files.values()).map((f) => new GoDirEntry(f));
    const dirs = Array.from(this.dirs.entries()).map(
      ([k, v]) => new GoDirEntry(new GoReadDirFile(k, v.entries()))
    );
    return Array.prototype.concat(files, dirs);
  }
}
