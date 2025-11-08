function addClass(cls, ...elems) {
  elems.forEach((e) => e.classList.add(cls));
}
function removeClass(cls, ...elems) {
  elems.forEach((e) => e.classList.remove(cls));
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
function resolve(p) {
  parts = p.split(".");
  v = window;
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
    b.addEventListener("click", function () {
      console.log("click", b);
      if (b.classList.contains("is-selected")) {
        target.setAttribute("hidden", "true");
        b.classList.remove("is-selected");
        b.classList.remove("is-outlined");
        b.classList.remove("is-primary");
      } else {
        target.removeAttribute("hidden");
        b.classList.add("is-selected");
        b.classList.add("is-outlined");
        b.classList.add("is-primary");
      }
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
  return e.value || e.textContent;
}
function safeUInt(v) {
  n = Number(v);
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
