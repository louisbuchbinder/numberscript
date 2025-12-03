function addClass(cls, ...elems) {
  elems.forEach((e) => e.classList.add(cls));
}
function removeClass(cls, ...elems) {
  elems.forEach((e) => e.classList.remove(cls));
}
function toggleClass(cls, ...elems) {
  elems.forEach((e) => e.classList.toggle(cls));
}
function hide(...elems) {
  elems.forEach((e) => e.setAttribute("hidden", "true"));
}
function unhide(...elems) {
  elems.forEach((e) => e.removeAttribute("hidden"));
}
function uint8ArrayFromSpaceSeparatedString(s) {
  const bytes = s.split(" ").map(Number);
  for (b of bytes) {
    if (Number.isNaN(b) || b < 0 || 255 < b) {
      throw new Error(
        "Unable to parse bytes. Expected space separated uint8, but instead got: " +
          s
      );
    }
  }
  return Uint8Array.from(bytes);
}
function arrToSpaceSeparatedString(b) {
  return b.map(String).join(" ");
}
function resolve(p, v = window) {
  parts = p.split(".");
  while (parts.length > 0) {
    v = v[parts.shift()];
  }
  return v;
}
function respondToVisibility(element, callback) {
  const options = {
    root: document.documentElement,
  };

  const observer = new IntersectionObserver((entries, observer) => {
    entries.forEach((entry) => {
      callback(entry.intersectionRatio > 0);
    });
  }, options);

  observer.observe(element);
}
document.addEventListener("DOMContentLoaded", function () {
  [].slice
    .call(document.getElementsByClassName("is-copy-button"))
    .forEach((b) => {
      const target = b.parentElement.querySelector(".is-copy-target");
      b.addEventListener("click", () => {
        navigator.clipboard.writeText(dataValue(target));
      });
    });

  [].slice.call(document.querySelectorAll(".is-menu-button")).forEach((b) => {
    const target = b.parentElement.querySelector(".menu-list");
    const icon = b.parentElement.querySelector(".icon > i");
    b.addEventListener("click", function () {
      if (b.classList.contains("is-selected")) {
        target.setAttribute("hidden", "true");
        icon.classList.remove("fa-angle-down");
        icon.classList.add("fa-angle-up");
      } else {
        target.removeAttribute("hidden");
        icon.classList.remove("fa-angle-up");
        icon.classList.add("fa-angle-down");
      }
      b.classList.toggle("is-selected");
      b.classList.toggle("is-outlined");
      b.classList.toggle("is-primary");
    });
  });

  [].slice
    .call(document.querySelectorAll("textarea.is-auto-resize"))
    .forEach((e) => {
      const resize = () => {
        e.style.height = "auto";
        e.style.height = e.scrollHeight + "px";
      };
      e.addEventListener("input", resize);
      respondToVisibility(e, resize);
      resize();
    });
});
function dataValue(e) {
  if (
    e.tagName.toLowerCase() === "input" &&
    e.getAttribute("type") === "file"
  ) {
    return e.files;
  }
  return e.value || e.textContent;
}
function firstGoFile(files) {
  if (files.length < 1) {
    throw new Error("missing expected input file");
  }
  return new GoFile(files[0]);
}
function newGoFS(files) {
  if (files.length < 1) {
    throw new Error("missing expected input file");
  }
  return new GoFS(files);
}
async function createOpfsFile(name) {
  return await newOpfsFile(name, { create: true });
}
async function newOpfsFile(name, opts) {
  if (name.length === 0) {
    throw new Error("missing expected file name");
  }
  return await OpfsFile.open(name, opts);
}
function safeUInt(v) {
  const n = Number(v);
  if (isNaN(n)) {
    throw new Error("expected uint, but instead got: " + v);
  }
  if (n < 0) {
    throw new Error("expected uint, but instead got: " + n);
  }
  if (!Number.isInteger(n)) {
    throw new Error("expected uint, but instead got: " + n);
  }
  if (!Number.isSafeInteger(n)) {
    throw new Error("integer out of range: " + n);
  }
  return n;
}
function safeInt(v) {
  const n = Number(v);
  if (isNaN(n)) {
    throw new Error("expected int, but instead got: " + v);
  }
  if (!Number.isInteger(n)) {
    throw new Error("expected int, but instead got: " + n);
  }
  if (!Number.isSafeInteger(n)) {
    throw new Error("integer out of range: " + n);
  }
  return n;
}
function safeFloat(v) {
  const n = Number(v);
  if (isNaN(n)) {
    throw new Error("expected float64, but instead got: " + v);
  }
  if (n < -1 * Number.MAX_VALUE || Number.MAX_VALUE < n) {
    throw new Error(
      "value must be within float64 range, but instead got: " + v
    );
  }
  return n;
}
function safeFloat32(v) {
  const maxFloat32 = 3.4028235 * Math.pow(10, 38);
  const n = Number(v);
  if (isNaN(n)) {
    throw new Error("expected float32, but instead got: " + v);
  }

  if (n < -1 * maxFloat32 || maxFloat32 < n) {
    throw new Error(
      "value must be within float32 range, but instead got: " + v
    );
  }
  return n;
}
