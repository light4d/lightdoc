import marked from 'marked'
import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

const state = {
  Path: '',
  readmeContent: ''
}
const mutations = {
  changeValue(state, newValue) {
    state.Path = newValue
    axios.get('http://localhost:8000' + state.Path).then((res) => {
      state.readmeContent = marked(res.data)
      console.log(res)
    })
  }
}
export default new Vuex.Store({
  state,
  mutations
})
