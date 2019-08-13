import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const state = {
  Node: {},
  readmeContent: '',
  number: '',
  resize: '',
  splitbar: false,
  fixed: false,
  sidebarMarginLeft: 'init'
}
export default new Vuex.Store({
  state
})
