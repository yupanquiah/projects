// @ts-check
import tailwindcss from '@tailwindcss/vite';
import { defineConfig, fontProviders } from 'astro/config';

// https://astro.build/config
export default defineConfig({
  vite: {
    plugins: [tailwindcss()]
  },
  experimental: {
    fonts: [
      {
        name: 'Poppins',
        cssVariable: '--font-poppins',
        provider: fontProviders.fontsource(),
        weights: [400, 500, 600, 700, 800],
        styles: ['normal'],
        subsets: ['latin']
      },
      {
        name: 'Volkhov',
        cssVariable: '--font-volkhov',
        provider: fontProviders.fontsource(),
        weights: [400, 700],
        styles: ['normal'],
        subsets: ['latin']
      }
    ]
  },
  site: 'https://lodging-cocha.com',
});
