import {createApp, watch} from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'




const app = createApp(App)

import * as _app from '../wailsjs/go/main/App'
import * as _runtime from '../wailsjs/runtime/runtime.js'
import _ from 'lodash';
import * as store from "./stores/store.js";
import * as utils from './scripts/utils.js'

app.config.globalProperties.app = _app;
app.config.globalProperties.runtime = _runtime;
app.config.globalProperties._ = _;
app.config.globalProperties.utils = utils;
app.config.globalProperties.store = store;


//import Navigation from "./components/Navigation/Navigation.vue";
//import Tooltip from "./components/Tooltip.vue";
import TextBox from "./components/Elements/TextBox.vue";
import Draggable from "./components/Draggable/Draggable.vue";
import ContextMenu from './components/ContextMenu/ContextMenu.vue'
import TreeView from "./components/TreeView/TreeView.vue";
import Button from "./components/Elements/Button.vue";
import Checkbox from "./components/Elements/Checkbox.vue";
import Table from "./components/Table/Table.vue";
import Dropdown from "./components/Elements/Dropdown.vue";
import Container from "./components/Elements/Container.vue";
import ToggleButton from "./components/Elements/ToggleButton.vue";

app
  .component('TextBox', TextBox)
  .component('Button', Button)
  .component('ContextMenu', ContextMenu)
  .component('Dropdown', Dropdown)
  .component('Draggable', Draggable)
  .component('Checkbox', Checkbox)
  .component('Table', Table)
  .component('TreeView', TreeView)
  .component('Container', Container)
  .component('ToggleButton', ToggleButton)

app.use(createPinia())
app.mount('#app')

