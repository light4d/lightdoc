import marked from 'marked'
import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

const state = {
  Path: '',
  readmeContent: '',
  number:200
}
const mutations = {
  changeValue(state, newValue) {
    state.Path = newValue
   
    if (state.Path.Nodes.length>0) {
      console.log(state,'rademe')
      axios.get('http://localhost:8000' + state.Path.Path + '/Readme.md ')
        .then((res) => {
          state.readmeContent = marked(res.data)
          console.log(res)
        })
        .catch((err)=>{
          console.log(111)
        })
    }else{
      console.log(state,'pdf')
      if(state.Path.Path.indexOf('.pdf')>-1){
        console.log('这是pdf文件')
        window.open('http://localhost:8000'+state.Path.Path)
        return
      }
      // if(state.Path.Path.indexOf('.md')>-1){
      //   axios.get('http://192.168.2.23:8000' + state.Path.Path)
      //   .then((res) => {
      //     console.log(res,'pdf')
          
      //     state.readmeContent = marked(res.data)
      //   })
      //   .catch((err)=>{
          
      //     console.log('errStort')
      //   })
      // }
      
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
