<template>
  <!-- @click.stop可以阻止事件冒泡 -->
  <div class="m-container" @click.stop="handleMobile">
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
  //在网页创建前根据路由获取内容
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
        console.log(err);
      });
  },
  methods: {
    handleMobile() {
      // 需要注意的是，手机端和电脑端的事件不一样，手机端因为没有鼠标，也就无法判断鼠标位置实现菜单栏的收展，所以在手机端需要点击内容部分实现菜单栏隐藏
      // 只有当菜单栏没有固定时，点击内容部分，菜单栏隐藏
      if (this.$store.state.fixed === false) this.$emit("handleMobile", false);
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
  padding: 0 20px;
  width: auto;
  transition: all 0.2s ease;
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
    transition: all 0.2s ease;
    padding: 0 10px;
    width: auto;
  }
}
</style>
