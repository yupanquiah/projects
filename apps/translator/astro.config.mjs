// @ts-check
import { defineConfig, fontProviders } from "astro/config";

// https://astro.build/config
export default defineConfig({
  fonts: [
    {
      provider: fontProviders.local(),
      name: "Geist Sans",
      cssVariable: "--geist-sans",
      options: {
        variants: [
          {
            weight: "100 900",
            style: "normal",
            display: "swap",
            src: ["./src/assets/fonts/GeistV.woff2"],
          },
        ],
      },
    },
  ],
});
