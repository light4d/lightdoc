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
  computed: {},
  methods: {
    changeContent(e) {
      //console.log("Search component: " + this.searchContent);
      setTimeout(() => {
        this.$emit("changeContent", this.searchContent);
      }, 400);
    }
  },
  mounted() {}
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
