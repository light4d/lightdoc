import marked from 'marked'
import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

const state = {
  Node: {},
  readmeContent: '',
  number: '',
  resize: ''
}
const mutations = {
  changeValue (state, newValue) {
    state.Node = newValue
    console.log(state.Path)

    var reg = /[^<>/\\\|:""\*\?]+\.\w+$/
    var matchs = state.Node.Path.match(reg)
    if (matchs != null) {
      console.log(matchs[0])
      state.Node.Name = state.Node.Path.match(reg)[0]
    }

    if (state.Node.Nodes.length > 0) {
      for (let i = 0; i < state.Node.Nodes.length; i++) {
        if (state.Node.Nodes[i].Node.indexOf('.md') == -1) {
          return
        } else {
          console.log(state.Node.Nodes[i].Node.match(reg))
        }
      }
      axios({
        method: 'FILE',
        url: state.Node.Path + '/README.md ',
        data: null
      }).then((res) => {
        state.readmeContent = marked(res.data)
        console.log(res)
      })
        .catch((err) => {
          console.log(err)
        })
    } else {
      console.log(state, 'pdf')
      if (state.Path.Node.indexOf('.pdf') > -1) {
        console.log('这是pdf文件')
        window.open('/' + state.Node.Path)
        return
      }
      if (state.Node.Path.indexOf('.md') > -1) {
        axios({
          method: 'FILE',
          url: state.Node.Path,
          data: null
        })
          .then((res) => {
            console.log(res, 'md')

            state.readmeContent = marked(res.data)
          })
          .catch((err) => {
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
