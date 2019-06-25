<template>
  <div class="m-container" v-html="this.$store.state.readmeContent"></div>
</template>

<script>
import marked from "marked";
import hljs from "highlight.js";

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
@import '../../static/monokai-sublime.css';

.m-container {
  padding: 0 30px 0 30px;
  overflow: auto;
  width: 100%;
}
</style>
