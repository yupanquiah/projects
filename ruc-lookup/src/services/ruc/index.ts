import { consultarRuc } from "./scraper";

console.log("Iniciando...");

try {
  const data = await consultarRuc("20114839176");
  console.log("Resultado:", JSON.stringify(data, null, 2));
} catch (e) {
  console.error("Error:", e);
}
