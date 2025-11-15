class GoFile {
  file: File;
  parts: Blob[];
  error: null | Error;

  constructor(f: File, limit: number = 100) {
    this.file = f;
    this.parts = [];
    for (let i = 0; i < f.size; i += limit) {
      this.parts.push(this.file.slice(i, i + limit));
    }
    this.error = null;
  }

  stat(): GoFileInfo {
    return new GoFileInfo(this.file);
  }

  async read(): Promise<Uint8Array<ArrayBufferLike>> {
    if (this.parts.length === 0) {
      return null;
    }

    const part = this.parts.shift();

    if (this.error != null) {
      throw this.error;
    }

    return part
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
