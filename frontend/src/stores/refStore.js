import {ref, computed, reactive, watch, onMounted, nextTick} from 'vue';
import {defineStore} from 'pinia';
import _ from 'lodash';

import * as app from '../../wailsjs/go/main/App';
import * as runtime from '../../wailsjs/runtime/runtime.js';
import * as utils from './../scripts/utils.js';

import {storeStore} from './store.js';

export const refStore = defineStore('ref', {
	state: () => {
		return {
			main: null,
			topMenu: null,
			mainTopMenu: null,
			topMenuGameData: null,
			sideMenu: null,
			container: null,
			content: null,
		};
	},
	getters: {},
	actions: {
		setContentPosition() {
			let rect = this.topMenu?.getBoundingClientRect();
			if (!rect) return;
			//this.content.style.paddingTop = `${rect.height}px`	//refs.content.style.marginTop = `${rect.height}px`
		},
		sidemenuResize(e) {
			this.container.style.marginLeft = `${e}px`;
			storeStore().state.sidemenuWidth = e;
		},
		async resizeWindowToFit() {
			let rectTopMenu = refStore().topMenu?.getBoundingClientRect();
			let rectContent = refStore().content?.getBoundingClientRect();
			let newSize = {
				x: parseInt(rectContent.width + storeStore().state.sidemenuWidth + 30),
				y: parseInt(rectContent.height + rectTopMenu.height + 60),
			};
			//app.WindowAction({opt: 'WindowSetSize', ...newSize})

			let currentSize = await runtime.WindowGetSize();
			let currentPos = await runtime.WindowGetPosition();

			let monitors = await runtime.ScreenGetAll();
			let currentMonitor = {i: null, isCurrent: null, isPrimary: null, width: null, height: null};
			for (let i = 0; i < monitors.length; i++) {
				if (monitors[i].isCurrent) currentMonitor = {i, ...monitors[i]};
			}

			console.log('resizeWindowToFit1', JSON.parse(JSON.stringify({newSize, currentSize, currentPos, currentMonitor, monitors})), currentPos.y + newSize.y, currentMonitor.height, '-->', currentMonitor.height - currentPos.y);
			//let spaceLeft = currentMonitor.width - currentSize.w
			if (currentPos.x + newSize.x > currentMonitor.width) console.log('!!!!!! w'); //newSize.w = currentMonitor.width - currentPos.x
			if (currentPos.y + newSize.y > currentMonitor.height) console.log('!!!!!! h'); //newSize.h = currentMonitor.height - currentPos.y

			if (currentPos.x + newSize.x > currentMonitor.width) newSize.x = Math.abs(currentMonitor.width - currentPos.x);
			if (currentPos.y + newSize.y + 400 > currentMonitor.height) newSize.y = Math.abs(currentMonitor.height - currentPos.y - 400);
			console.log('resizeWindowToFit2', {newSize});
			await nextTick();
			await app.WindowAction({opt: 'WindowSetSize', ...newSize});
			//runtime.WindowSetSize(newSize)
			//let environment = await runtime.Environment()
			/* 
			WindowGetSize()
			// JS: WindowSetSize(size: Size)
			// JS: WindowGetPosition() : Position
			// JS: WindowSetPosition(position: Position)
			// JS: WindowSetTitle(title: string)
			// JS: WindowSetAlwaysOnTop(b: Boolen)
			
			JS: WindowIsNormal() bool	// Returns true if the window not minimised, maximised or fullscreen



			// JS: WindowFullscreen()
			JS: WindowUnmaximise()
			JS: WindowIsMaximised() bool
			JS: WindowToggleMaximise()
			
			// JS: WindowMinimise()
			JS: WindowUnminimise()
			JS: WindowIsMinimised() bool
			
			JS: WindowSetBackgroundColour(R, G, B, A)




			// JS: WindowHide()
			// JS: WindowShow()
			// JS: WindowCenter()
			*/
		},
	},
});
