<template>
  <div class="z-container">
    <!-- 右上角的样式选择 -->
    <div class="styleSel">
      <el-select
        v-model="value"
        size="mini"
        placeholder="样式选择"
        style="width: 100px"
        @change="getValue"
      >
        <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item"></el-option>
      </el-select>
    </div>

    <!-- 侧边栏 -->
    <div id="sidebar" :style="{
        position: 'fixed'
      }">
      <v-aside :mobileV="mobileV" @handleMobile="handleMobile"></v-aside>
      <transition name="el-zoom-in-center">
        <div id="split-bar" @mousedown="resize" v-show="mobileV"></div>
      </transition>
    </div>

    <!-- 内容区 -->
    <v-markdown
      id="main"
      :style="{
        marginLeft: this.$store.state.splitbar ? this.$store.state.sidebarMarginLeft === 'init' ? '250px' : this.$store.state.sidebarMarginLeft+'px': '20px'
      }"
      :mobileV="mobileV"
      @handleMobile="handleMobile"
    ></v-markdown>
  </div>
</template>
<script>
import vAside from "../components/aside.vue";
import vMarkdown from "../components/markdown.vue";
import smShow from "../components/smShow.vue";
import $ from "jquery";

export default {
  data() {
    return {
      options: [
        {
          value: "1",
          label: "Green",
          rgb: "#009688"
        },
        {
          value: "2",
          label: "Blue",
          rgb: "#409EFF"
        },
        {
          value: "3",
          label: "Black",
          rgb: "#303133"
        }
      ],
      value: "",
      mobileV: false
    };
  },
  components: {
    vAside,
    vMarkdown,
    smShow
    //BackToTop
  },
  methods: {
    // 移动端适用，通过给value赋值人为改变侧边栏收展
    handleMobile(value) {
      if (value !== undefined) {
        this.mobileV = value;
      } else {
        this.mobileV = !this.mobileV;
      }
      if (this.mobileV === true)
        document.querySelector("#sidebar").style.backgroundColor = "white";
      else
        document.querySelector("#sidebar").style.backgroundColor =
          "transparent";
    },
    getValue(color) {
      var elements = [];
      // 获取所有h1元素列表
      elements[0] = document.querySelectorAll("h1");
      elements[1] = document.querySelectorAll("h2");
      elements[2] = document.querySelectorAll("h3");
      elements[3] = document.querySelectorAll("h4");
      for (var i = 0; i < elements.length; i++) {
        for (var n = 0; n < elements[i].length; n++) {
          elements[i][n].style.color = color.rgb;
        }
      }
    },
    resize(e) {
      e.preventDefault();
      var min = 200;
      var max = 800;
      var mainmin = 200;
      var that = this;
      // 通过调整margin-left改变内容区和侧边栏宽度 侧边栏宽度就是内容区的margin-left
      $(document).mousemove(function(e) {
        e.preventDefault();
        var x = e.pageX - $("#sidebar").offset().left;
        if (x > min && x < max && e.pageX < $(window).width() - mainmin) {
          $("#sidebar").css("width", x);
          $(".aside-container").css("width", x - 5);
          $("#main").css("margin-left", x);
          that.$store.state.sidebarMarginLeft = x;
        }
      });
    }
  },
  mounted() {
    //取消resize方法中在document中添加的mousemove方法
    $(document).mouseup(function(e) {
      $(document).unbind("mousemove");
    });
  }
};
</script>
<style lang="stylus" scoped>
.z-container {
  width: 100%;
  position: relative;
  overflow: auto;
  height: 100%;
}

.styleSel {
  position: fixed;
  right: 10px;
  top: 10px;
}

#sidebar {
  width: 250px;
  height: 100%;
  float: left;
  display: flex;
  flex-direction: column;
  flex-wrap: wrap;
  transition: opacity 0.2s ease;
}

#split-bar {
  background-color: #ccc;
  height: 100%;
  width: 6px;
  cursor: col-resize;
  float: right;
  align-self: flex-end;
  z-index: 1000;
}

#main {
  margin-left: 20px;
  height: 100%;
}

@media screen and (max-width: 768px) {
  #sidebar {
    display: flex;
    flex-direction: column;
    width: 70%;
    overflow: hidden;
    height: 100%;
    transition: opacity 0.3s ease;
  }

  #split-bar {
    position: absolute;
  }
}
</style>
