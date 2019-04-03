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
    if (state.Path.Nodes.length>0) {
      axios.get('/api' + state.Path.Path+'/Readme.md ' )
        .then((res) => {
          state.readmeContent = marked(res.data)
          console.log(res)
        })
        .catch((err)=>{
          console.log(111)
        })
    }else{
       axios.get('/api' + state.Path.Path)
         .then((res) => {
           state.readmeContent = marked(res.data)
           console.log(res,'xia')
         })
    }
    // axios.get('/api' + state.Path)
    // .then((res) => {
    //   state.readmeContent = marked(res.data)
    //   console.log(res)
    // })
    // .catch((err)=>{
    //   alert(err)
    // })
  }
}
export default new Vuex.Store({
  state,
  mutations
})
