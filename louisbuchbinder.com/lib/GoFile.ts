class GoFile {
  file: File;
  generator: Generator<Blob, null, undefined>;
  done: boolean;
  error: null | Error;

  static ThirtyTwoKB = Math.pow(2, 10) * 32;
  static OneMB = Math.pow(2, 20);
  static ThirtyTwoMB = Math.pow(2, 20) * 32;
  static FiveHundredMB = Math.pow(2, 20) * 500;
  static FiveGB = Math.pow(2, 30) * 5;

  static BufferSize(size: number): number {
    if (size > GoFile.FiveGB) {
      return this.ThirtyTwoMB;
    }

    if (size > GoFile.FiveHundredMB) {
      return this.OneMB;
    }

    return this.ThirtyTwoKB;
  }

  static *Generator(f: File): Generator<Blob, null, undefined> {
    let i = 0;
    const limit = GoFile.BufferSize(f.size);
    while (i < f.size) {
      yield f.slice(i, i + limit);
      i += limit;
    }
    return null;
  }

  constructor(f: File) {
    this.file = f;
    this.generator = GoFile.Generator(f);
    this.done = false;
    this.error = null;
  }

  stat(): GoFileInfo {
    return new GoFileInfo(this.file);
  }

  async read(): Promise<Uint8Array<ArrayBufferLike>> {
    if (this.done) {
      this.error =
        this.error ||
        new Error("unexpected call to read after processing the complete file");
    }
    if (this.error !== null) {
      throw this.error;
    }

    const next = this.generator.next();

    if (next.done) {
      this.done = true;
      return next.value;
    }

    return next.value
      .arrayBuffer()
      .then((ab) => {
        const arr = new Uint8Array(ab);
        return arr;
      })
      .catch((err) => {
        this.error = err;
        throw err;
      });
  }

  async close(): Promise<null> {
    return null;
  }
}

class GoFileInfo {
  file: File;
  constructor(f: File) {
    this.file = f;
  }

  name(): string {
    return this.file.name;
  }

  mode(): null {
    return null;
  }

  size(): number {
    return this.file.size;
  }

  modTime(): number {
    return this.file.lastModified;
  }

  isDir(): boolean {
    return false;
  }

  sys(): null {
    return null;
  }
}
