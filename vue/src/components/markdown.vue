<template>
  <div class="m-container" @click="handleMobile">
    <el-table
      class="load"
      v-loading="isLoading"
      element-loading-text="拼命加载中"
      element-loading-spinner="el-icon-loading"
      element-loading-background="#DCDFE6"
      style="width: 100%"
    ></el-table>
    <div id="content" v-html="this.$store.state.readmeContent"></div>
  </div>
</template>

<script>
import marked from "marked";
import hljs from "highlight.js";

export default {
  name: "vMarkdown",
  data() {
    return {
      isLoading: false,
      timer: ""
    };
  },
  props: {
    mobileV: Boolean
  },
  watch: {
    $route(to, from) {
      console.log(to);
      this.isLoading = true;
      this.axios({
        method: "GET",
        url: to.path,
        data: null
      })
        .then(res => {
          this.$store.state.readmeContent = marked(res.data);
          this.isLoading = false;
        })
        .catch(err => {
          console.log(111);
        });
      this.timer = setTimeout(this.apiFailed, 3000);
    }
  },
  created() {
    this.isLoading = true;
    this.axios({
      method: "GET",
      url: "/" + this.$route.params.pageId,
      data: null
    })
      .then(res => {
        this.$store.state.readmeContent = marked(res.data);
        this.isLoading = false;
      })
      .catch(err => {
        console.log(111);
      });
    this.timer = setTimeout(this.apiFailed, 8000);
    // console.log(this.$store.state.readmeContent);
    // marked.setOptions({
    //   renderer: new marked.Renderer(),
    //   highlight: function(readmeContent) {
    //     return hljs.highlightAuto(readmeContent).value;
    //   },
    //   pedantic: false,
    //   gfm: true,
    //   tables: true,
    //   breaks: false,
    //   sanitize: false,
    //   smartLists: true,
    //   smartypants: false,
    //   xhtml: false
    // });
  },
  methods: {
    handleMobile() {
      this.$emit("handleMobile", false);
    },
    changeData(value, render) {
      console.log(render);
    },
    apiFailed() {
      this.isLoading = false;
    }
  }
};
</script>

<style lang="stylus" scoped>
// @import '../../static/monokai-sublime.css';
.m-container {
  display: flex;
  padding: 0 10px;
  width: auto;
}

.alertMsg {
  position: absolute;
  width: auto;
  top: 10%;
  left: 50%;
  transform: translateX(40%);
}

#content {
  width: 100%;
}

@media screen and (max-width: 768px) {
  .m-container {
    margin-left: 10px !important;
  }
}
</style>
