<template>
  <div class="m-container">
    <div v-html="this.$store.state.readmeContent"></div>
  </div>
</template>

<script>
import marked from "marked";
import hljs from "highlight.js";

export default {
  name: "vMarkdown",
  data() {
    return {
      defaultData: "preview",
      isLoading: false
    };
  },
  watch: {
    $route(to, from) {
      this.axios({
        method: "GET",
        url: to.path,
        data: null
      })
        .then(res => {
          this.$store.state.readmeContent = marked(res.data);
        })
        .catch(err => {
          console.log(111);
        });
    }
  },
  created() {
    this.axios({
      method: "GET",
      url: "/" + this.$route.params.pageId,
      data: null
    })
      .then(res => {
        this.$store.state.readmeContent = marked(res.data);
      })
      .catch(err => {
        console.log(111);
      });
    console.log(this.$store.state.readmeContent);
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
  width: auto;
}
</style>
