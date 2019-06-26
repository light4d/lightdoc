<template>
  <div class="search">
    <input disabled class="inpt" v-model="searchContent" placeholder="搜索暂不可用！">
    <i class="el-icon-search iconSearch"></i>
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

  .inpt {
    width: 100%;
    text-indent: 40px;
    display: block;
    height: 30px;
    background: rgba(0, 0, 0, 0.1);
    // border-bottom:1px solid #ccc
    // border-top:1px solid #ccc
  }

  .iconSearch {
    position: absolute;
    left: 5%;
    top: 25%;
  }
}
</style>
