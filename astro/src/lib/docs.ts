import fs from "node:fs";
import path from "node:path";

const DOCS_DIR = path.resolve("../generated/docs");

export interface EndpointDoc {
  title: string;
  slug: string;
  category: string;
  endpoint: string;
  method: string;
  description: string;
}

export function getDocs(): EndpointDoc[] {
  const files = fs.readdirSync(DOCS_DIR);

  return files
    .map((file) => {
      const fullPath = path.join(DOCS_DIR, file);
      return JSON.parse(fs.readFileSync(fullPath, "utf-8"));
    })
    .sort((a, b) => a.title.localeCompare(b.title));
}

export function getCategories(docs: EndpointDoc[]): string[] {
  const cats = new Set(docs.map((d) => d.category));
  return [...cats].sort();
}

export function getDocsByCategory(
  docs: EndpointDoc[],
  category: string,
): EndpointDoc[] {
  return docs.filter((d) => d.category === category);
}
