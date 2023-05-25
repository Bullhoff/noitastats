import {defineConfig} from 'vite';
import vue from '@vitejs/plugin-vue';
//import {fileURLToPath} from 'url';

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [
		vue(),
		/* {
          template: {
            compilerOptions: {
              isCustomElement: (tag) => ['ContextMenu'].includes(tag),
            },
          },
        } */
	],
	/* resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),
    },
  }, */
  
});
