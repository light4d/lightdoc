<template>
  <div class="m-container">
    <back-to-top bottom="4em" right="4em">
      <button type="button" class="btn-to-top">
        <i class="el-icon-caret-top"></i>
      </button>
    </back-to-top>
    <div v-html="this.$store.state.readmeContent"></div>
  </div>
</template>

<script>
import marked from "marked";
import hljs from "highlight.js";
import BackToTop from "vue-backtotop";

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
    BackToTop
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
