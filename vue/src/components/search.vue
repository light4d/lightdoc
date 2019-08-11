<template>
  <div class="search">
    <el-input
      type="text"
      placeholder="请输入内容"
      clearable
      prefix-icon="el-icon-search"
      v-model="searchContent"
      @input="changeContent"
    ></el-input>
  </div>
</template>
<script>
export default {
  name: "vSearch",
  data() {
    return {
      searchContent: "text"
    };
  },
  methods: {
    changeContent(e) {
      //在输入停止0.4秒后再进行搜索
      setTimeout(() => {
        this.$emit("changeContent", this.searchContent);
      }, 400);
    }
  }
};
</script>

<style lang="stylus" scoped>
.search {
  position: relative;
  width: 90%;
  left: 5%;
  right: 5%;
}
</style>
