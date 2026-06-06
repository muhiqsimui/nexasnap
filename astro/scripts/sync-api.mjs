import fs from "node:fs";
import path from "node:path";

const source = path.resolve("../generated/api");
const target = path.resolve("./public/api");

fs.rmSync(target, {
  recursive: true,
  force: true,
});

fs.cpSync(source, target, {
  recursive: true,
});

console.log("API synced");
