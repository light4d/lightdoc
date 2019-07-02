<template>
  <div class="search">
    <el-input
      type="text"
      placeholder="请输入内容"
      clearable
      prefix-icon="el-icon-search"
      v-model="searchContent"
    ></el-input>
  </div>
</template>
<script>
export default {
  name: "vSearch",
  data() {
    return {
      searchContent: ""
    };
  },
  computed: {},
  methods: {},
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
