<template>
	<div class="container" v-html="this.$store.state.readmeContent">	
	</div>
</template>

<script>
import marked from 'marked'
import hljs from "highlight.js"
import javascript from 'highlight.js/lib/languages/javascript'
export default {
	name:'vMarkdown',
	props:{
		Path:{
			type:Array
		}
	},
	data() {
        return {//value的值是经过markdown解析后的文本，可使用`@change="changeData"`在控制台打印显示
            defaultData: "preview",
			readmeContent:'',
			pathmd:''
        };
    },
	computed:{
// 		this.axios.get('/api'+this.Path).then((res)=>{
// 			this.readmeContent=marked(res.data)
// 		})
// 		marked.setOptions({
// 		  renderer: new marked.Renderer(),
// 		  highlight: function(readmeContent) {
// 		    return hljs.highlightAuto(readmeContent).value;
// 		  },
// 		  pedantic: false,
// 		  gfm: true,
// 		  tables: true,
// 		  breaks: false,
// 		  sanitize: false,
// 		  smartLists: true,
// 		  smartypants: false,
// 		  xhtml: false
// 		}
	},
	mounted(){
		console.log( this.Path,'-----')
		this.axios.get('http://localhost:8000'+this.Path).then((res)=>{
			this.readmeContent=marked(res.data)
		})
		marked.setOptions({
          renderer: new marked.Renderer(),
          highlight: function(readmeContent) {
            return hljs.highlightAuto(readmeContent).value;
          },
          pedantic: false,
          gfm: true,
          tables: true,
          breaks: false,
          sanitize: false,
          smartLists: true,
          smartypants: false,
          xhtml: false
        }
      );
	},
	watch: {
		info() {
		  this.pathmd = this.Path
		}
	},
    methods: {
        changeData(value, render) {
            console.log(render);
        }
    }
}
</script>

<style lang="stylus" scoped>
	@import '../assets/monokai-sublime.css'
	.container
		padding:20px 30px
</style>
