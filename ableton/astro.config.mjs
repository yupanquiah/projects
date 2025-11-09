// @ts-check
import { defineConfig } from "astro/config";

// https://astro.build/config
export default defineConfig({
  experimental: {
    fonts: [
      {
        provider: "local",
        name: "futura-pt",
        cssVariable: "--font-futura-pt",
        variants: [
          {
            weight: "normal",
            style: "normal",
            src: [
              "src/assets/fonts/futura-pt.woff2",
              "src/assets/fonts/futura-pt.woff",
            ],
            display: "swap",
          },
          {
            weight: "bold",
            style: "normal",
            src: [
              "src/assets/fonts/futura-pt-medium.woff2",
              "src/assets/fonts/futura-pt-medium.woff",
            ],
            display: "swap",
          },
          {
            weight: "normal",
            style: "italic",
            src: [
              "src/assets/fonts/futura-pt-oblique.woff2",
              "src/assets/fonts/futura-pt-oblique.woff",
            ],
            display: "swap",
          },
        ],
      },
    ],
  },
});
