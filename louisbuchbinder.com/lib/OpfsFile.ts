interface IGoWriteCloser {
  write(p: Uint8Array<ArrayBuffer>): Promise<number>;
  close(): Promise<void>;
}

class OpfsFile extends GoFile implements IGoWriteCloser, IGoFile {
  fileHandle: FileSystemFileHandle | null;
  writable: FileSystemWritableFileStream | null;
  position: number;
  closed: boolean;

  private constructor() {
    super(null);
    this.fileHandle = null;
    this.writable = null;
    this.position = 0;
    this.closed = false;
  }

  // Async factory: open (and create if needed) a file at `path` in the
  // Origin Private File System and return a writer.
  static async open(
    path: string,
    opts?: { create?: boolean }
  ): Promise<OpfsFile> {
    if (
      typeof navigator === "undefined" ||
      !navigator.storage ||
      !navigator.storage.getDirectory
    ) {
      throw new Error("OPFS not supported in this environment");
    }

    const create = opts && opts.create === true;

    // get the origin private file system root
    const root = await (navigator.storage as any).getDirectory();

    // Walk/create directories as necessary
    const parts = path.split("/").filter((p) => p.length > 0);
    let current: FileSystemDirectoryHandle = root;
    for (let i = 0; i < parts.length - 1; i++) {
      const part = parts[i];
      // getDirectoryHandle may throw if not present and create=false
      current = await current.getDirectoryHandle(part, { create: create });
    }

    const fileName = parts.length > 0 ? parts[parts.length - 1] : path;
    // obtain the file handle
    let fileHandle: FileSystemFileHandle | null = null;
    try {
      fileHandle = await current.getFileHandle(fileName, {
        create: create,
      });
    } catch (err) {
      throw new Error("failed to obtain file handle: " + String(err));
    }

    return await OpfsFile.fromFileHandle(fileHandle);
  }

  static async fromFileHandle(
    fileHandle: FileSystemFileHandle
  ): Promise<OpfsFile> {
    const writer = new OpfsFile();
    writer.fileHandle = fileHandle;
    // create writable stream (deferred open of writable until first write)
    writer.position = 0;
    writer.writable = null;
    writer.closed = false;
    writer.file = (await writer.fileHandle.getFile()) as File;
    return writer;
  }

  // Low-level ensure writable stream is ready
  private async ensureWritable(): Promise<void> {
    if (this.closed) {
      throw new Error("writer already closed");
    }
    if (!this.fileHandle) {
      throw new Error("fileHandle not initialized");
    }
    if (!this.writable) {
      // createWritable may accept options; we use default which truncates file.
      // To append, we write sequentially and keep position.
      this.writable = await (this.fileHandle as any).createWritable();
    }
  }

  // Write bytes to the file. Returns the number of bytes written.
  async write(data: Uint8Array<ArrayBuffer>): Promise<number> {
    await this.ensureWritable();
    if (!this.writable) {
      throw new Error("writable stream not available");
    }

    try {
      // write sequentially
      await this.writable.write(data);
      this.position += data.byteLength;
      return data.byteLength;
    } catch (err) {
      this.closed = true;
      throw err;
    }
  }

  // Close the writer and flush data
  async close(): Promise<void> {
    if (this.closed) {
      return;
    }
    if (this.writable) {
      await this.writable.close();
      this.writable = null;
    }
    this.closed = true;
  }
}
