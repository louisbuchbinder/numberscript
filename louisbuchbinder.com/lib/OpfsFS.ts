interface IGoOpfsFS extends IGoFS {
  mkdirAll(path: string): Promise<null>;
  removeAll(path: string): Promise<null>;
}

class OpfsFS implements IGoOpfsFS {
  root: FileSystemDirectoryHandle;

  private constructor(root: FileSystemDirectoryHandle) {
    this.root = root;
  }

  static async fromRoot(path: string): Promise<OpfsFS> {
    if (
      typeof navigator === "undefined" ||
      !navigator.storage ||
      !navigator.storage.getDirectory
    ) {
      throw new Error("OPFS not supported in this environment");
    }
    const opfsRoot = await navigator.storage.getDirectory();
    const parts = path.split("/").filter(Boolean);
    let root = opfsRoot;
    while (parts.length > 0) {
      root = await root.getDirectoryHandle(parts.shift());
    }
    return new OpfsFS(root);
  }

  async open(p: string): Promise<IGoFile> {
    if (p === "") {
      return await newGoReadDirFile(this.root);
    }
    try {
      const parts = p.split("/").filter(Boolean);
      let h = this.root;
      while (parts.length > 0) {
        h = await h.getDirectoryHandle(parts.shift());
      }
      return await newGoReadDirFile(h);
    } catch (err) {
      if (!(err instanceof DOMException) || err.name !== "TypeMismatchError") {
        throw err;
      }
    }
    const fileHandle = await this.root.getFileHandle(p);
    return OpfsFile.fromFileHandle(fileHandle);
  }

  async mkdirAll(p: string): Promise<null> {
    let h = this.root;
    const parts = p.split("/").filter(Boolean);
    while (parts.length > 0) {
      h = await h.getDirectoryHandle(parts.shift(), { create: true });
    }
    return null;
  }

  async removeAll(p: string): Promise<null> {
    const [h, parent] = await getHandle(this.root, p);
    if (h instanceof FileSystemDirectoryHandle) {
      for await (const entry of h.entries()) {
        await h.removeEntry(entry[0], { recursive: true });
      }
      await (h as any).remove();
    } else if (h instanceof FileSystemFileHandle) {
      await parent.removeEntry(h.name);
    } else {
      throw new Error("unexpected FileSystemHandle type");
    }
    return null;
  }
}

async function getHandle(
  r: FileSystemDirectoryHandle,
  p: string
): Promise<[FileSystemHandle, FileSystemDirectoryHandle]> {
  let parent = null;
  let h = r;
  let parts = p.split("/").filter(Boolean);
  while (parts.length > 0) {
    const part = parts.shift();
    parent = h;
    try {
      h = await h.getDirectoryHandle(part);
    } catch (err) {
      if (
        !(err instanceof DOMException) ||
        err.name !== "TypeMismatchError" ||
        parts.length > 0
      ) {
        throw err;
      }
      return [await h.getFileHandle(part), parent];
    }
  }
  return [h, parent];
}

async function newGoReadDirFile(
  dirHandle: FileSystemDirectoryHandle
): Promise<GoReadDirFile> {
  const entries: IGoDirEntry[] = [];

  for await (const entry of dirHandle.entries()) {
    const entryHandle = entry[1] as FileSystemHandle;
    if (entryHandle.kind === "directory") {
      const entryDirHandle = entryHandle as FileSystemDirectoryHandle;
      const readDirFile = await newGoReadDirFile(entryDirHandle);
      entries.push(new GoDirEntry(await readDirFile.stat()));
    } else if (entryHandle.kind === "file") {
      const entryFileHandle = entryHandle as FileSystemFileHandle;
      const file = new GoFile(await entryFileHandle.getFile());
      const dirEntry = new GoDirEntry(await file.stat());
      entries.push(dirEntry);
    } else {
      throw new Error(
        "unexpected FileSystemHandle.kind, got: " + entryHandle.kind
      );
    }
  }
  return new GoReadDirFile(dirHandle.name, entries);
}
