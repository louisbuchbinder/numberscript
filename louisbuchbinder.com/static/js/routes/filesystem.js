var filesystemTemplate,
  filesystemFolderTreeFileTemplate,
  filesystemFolderTreeFolderTemplate,
  filesystemFileListFileTemplate,
  filesystemFileListFolderTemplate;

const filesystemReady = Promise.all([
  fetch("/html/filesystem.html")
    .then((r) => r.text())
    .then((t) => (filesystemTemplate = t)),
  fetch("/html/filesystem-foldertree-file.html")
    .then((r) => r.text())
    .then((t) => (filesystemFolderTreeFileTemplate = t)),
  fetch("/html/filesystem-foldertree-folder.html")
    .then((r) => r.text())
    .then((t) => (filesystemFolderTreeFolderTemplate = t)),
  fetch("/html/filesystem-filelist-file.html")
    .then((r) => r.text())
    .then((t) => (filesystemFileListFileTemplate = t)),
  fetch("/html/filesystem-filelist-folder.html")
    .then((r) => r.text())
    .then((t) => (filesystemFileListFolderTemplate = t)),
]);

function filesystemFolderTreeFile(dat) {
  return filesystemFolderTreeFileTemplate.replace("{{.Name}}", dat.name);
}
function filesystemFolderTreeFolder(dat) {
  return filesystemFolderTreeFolderTemplate.replace("{{.Name}}", dat.name);
}
function filesystemFileListFile(dat) {
  return filesystemFileListFileTemplate
    .replace("{{.Name}}", dat.name)
    .replace("{{.Type}}", dat.type)
    .replace("{{.Size}}", dat.size)
    .replace("{{.Modified}}", dat.modified);
}
function filesystemFileListFolder(dat) {
  return filesystemFileListFolderTemplate
    .replace("{{.Name}}", dat.name)
    .replace("{{.Type}}", dat.type)
    .replace("{{.Size}}", "-")
    .replace("{{.Modified}}", "-");
}

function filesystemFolderTree(FS) {
  return FS.children
    .map((dat) => {
      if (dat.type === "folder") {
        return filesystemFolderTreeFolder(dat);
      } else {
        return filesystemFolderTreeFile(dat);
      }
    })
    .join("\n");
}

function filesystemFileList(FS) {
  return FS.children
    .map((dat) => {
      if (dat.type === "folder") {
        return filesystemFileListFolder(dat);
      } else {
        return filesystemFileListFile(dat);
      }
    })
    .join("\n");
}

// async function filesystemFolderTreeHandler(req) {
//   await filesystemReady;
//   const url = new URL(req.url);
//   try {
//     const FS = JSON.parse(url.searchParams.get("FS"));
//     const content = filesystemFolderTree(FS);
//     return new Response(content, {
//       status: 200,
//       headers: { "Content-Type": "text/html; charset=utf-8" },
//     });
//   } catch (err) {
//     console.error(err);
//     return new Response("Bad request: " + err.message, {
//       status: 400,
//     });
//   }
// }

async function filesystemHandler(req) {
  await filesystemReady;
  const url = new URL(req.url);
  try {
    const FS = JSON.parse(url.searchParams.get("FS"));
    const folderTree = filesystemFolderTree(FS);
    const fileList = filesystemFileList(FS);
    return new Response(
      filesystemTemplate
        .replace("{{.FolderTree}}", folderTree)
        .replace("{{.FileList}}", fileList),
      {
        status: 200,
        headers: { "Content-Type": "text/html; charset=utf-8" },
      }
    );
  } catch (err) {
    console.error(err);
    return new Response("Bad request: " + err.message, {
      status: 400,
    });
  }
}
