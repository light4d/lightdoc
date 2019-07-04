<template>
  <div class="z-container">
    <div class="content">
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
      <div id="sidebar">
        <v-aside></v-aside>
        <div id="split-bar" v-if="this.$store.state.fixed"  @mousedown="resize"></div>
      </div>
      <v-markdown id="main"></v-markdown>
    </div>
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
      value: ""
    };
  },
  components: {
    vAside,
    vMarkdown,
    smShow
  },
  methods: {
    getValue(color) {
      var elements = [];
      elements[0] = document.querySelectorAll("h1");
      elements[1] = document.querySelectorAll("h2");
      elements[2] = document.querySelectorAll("h3");
      elements[3] = document.querySelectorAll("h4");
      console.log(elements);
      for (var i = 0; i < elements.length; i++) {
        for (var n = 0; n < elements[i].length; n++) {
          elements[i][n].style.color = color.rgb;
        }
      }
    },
    resize(e) {
      e.preventDefault();
      var min = 300;
      var max = 3600;
      var mainmin = 200;
      $(document).mousemove(function(e) {
        e.preventDefault();
        var x = e.pageX - $("#sidebar").offset().left;
        if (x > min && x < max && e.pageX < $(window).width() - mainmin) {
          $("#sidebar").css("width", x);
          $("#main").css("margin-left", x);
        }
      });
    }
  },
  mounted() {
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
  position: absolute;
  right: 10px;
  top: 10px;
}

#sidebar {
  //background-color: IndianRed;
  width: 320px;
  height 100%;
  float: left;
}

#split-bar {
  background-color: #ccc;
  height: 100%;
  width: 6px;
  cursor: col-resize;
  float: right;
  z-index 1000;
}
#main {
  margin-left: 200px;
  //background-color: BurlyWood;
  height:100%;
}
</style>
