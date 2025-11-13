class GoFile {
  file: File;
  parts: Blob[];
  index: number;
  error: null | Error;

  constructor(f: File, limit: number) {
    this.file = f;
    this.parts = [];
    for (let i = 0; i < f.size; i += limit) {
      this.parts.push(this.file.slice(i, i + limit));
    }
    this.index = this.parts.length > 0 ? 0 : -1;
    this.error = null;
  }

  stat(): GoFileInfo {
    return new GoFileInfo(this.file);
  }

  async read(): Promise<Uint8Array<ArrayBufferLike>> {
    if (this.index < 0 || this.index >= this.parts.length) {
      return null;
    }

    const part = this.parts[this.index];

    if (this.error != null) {
      throw this.error;
    }

    return part
      .arrayBuffer()
      .then((ab) => {
        this.index++;
        return new Uint8Array(ab);
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
