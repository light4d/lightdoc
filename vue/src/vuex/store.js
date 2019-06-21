import marked from 'marked'
import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'
 
Vue.use(Vuex)

const state = {
  Path: '',
  readmeContent: '',
  number:'',
  resize:''
}
const mutations = {
  changeValue(state, newValue) {
    state.Path = newValue
    console.log(state.Path)

    var reg = /[^<>/\\\|:""\*\?]+\.\w+$/
    var matchs=state.Path.Path.match(reg)
    if(matchs!=null){
      console.log(matchs[0])
      state.Path.Name = state.Path.Path.match(reg)[0]
    }
   
    
    if (state.Path.Nodes.length>0) {
      for(let i=0;i<state.Path.Nodes.length;i++){
        if(state.Path.Nodes[i].Path.indexOf('.md')==-1){
          return
        }else{
          console.log(state.Path.Nodes[i].Path.match(reg))
        }
      }
      axios({
				method: 'FILE',
				url: "/" + state.Path.Path + '/README.md ',
				data:null,
			}).then((res) => {
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
        window.open("/"+state.Path.Path)
        return
      }
      if(state.Path.Path.indexOf('.md')>-1){
        axios({
          method: 'FILE',
          url: "/"+ state.Path.Path,
          data:null,
        })
        .then((res) => {
          console.log(res,'md')
          
          state.readmeContent = marked(res.data)
        })
        .catch((err)=>{
          
          console.log('errStort')
        })
      }
      
    }

  }
}
export default new Vuex.Store({
  state,
  mutations
})
