<template>
  <div class="m-container">
    <h1>我们是最棒的</h1>
    <h2>Come on!</h2>
    <h3>lets do this</h3>
    <h2>yeah</h2>
    <vm-back-top :bottom="100" :duration="1000" :timing="'ease'"></vm-back-top>
    <div v-html="this.$store.state.readmeContent"></div>
  </div>
</template>

<script>
import marked from "marked";
import hljs from "highlight.js";
import Vue from "vue";
import VmBackTop from "vue-multiple-back-top";
import "element-ui/lib/theme-chalk/display.css";

export default {
  name: "vMarkdown",
  props: {
    Path: {
      type: Array
    }
  },
  data() {
    return {
      defaultData: "preview"
    };
  },
  components: {
    VmBackTop
  },
  mounted() {
    this.axios({
      method: "FILE",
      url: `/README.md`,
      data: null
    })
      .then(res => {
        this.$store.state.readmeContent = marked(res.data);
      })
      .catch(err => {
        console.log(111);
      });
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
    });
  },
  watch: {
    info() {
      this.pathmd = this.Path;
    }
  },
  methods: {
    changeData(value, render) {
      console.log(render);
    }
  }
};
</script>

<style lang="stylus" scoped>
// @import '../../static/monokai-sublime.css';
.m-container {
  padding: 0 30px 0 30px;
  overflow: auto;
  width: 100%;
}
</style>
