<template>
  <div>
    <div class="menu-con" ref="menu" @mouseover="handleAside" v-show="OppositeMobileV">菜单</div>
    <div
      class="aside-container"
      ref="aside"
      @mouseleave="handleHide"
      v-show="mobileV"
      :style="{
					display:'none',}"
    >
      <v-header @fixed="handleFixed"></v-header>
      <v-search></v-search>
      <div class="tree">
        <el-tree
          :data="data"
          :props="defaultProps"
          @node-click="handleNodeClick"
          id="el-tree"
          :render-content="renderContent"
        ></el-tree>
      </div>
    </div>
  </div>
</template>

<script>
import vSearch from "./search.vue";
import vHeader from "./header.vue";

export default {
  data() {
    return {
      data: [],
      defaultProps: {
        children: "Nodes",
        label: "Name",
        className: ""
      },
      W: 0,
      resize: "",
      showdiv: false, // 控制鼠标显示样式
      readmeContent: {},
      fixed: false
    };
  },
  props: {
    mobileV: Boolean
  },
  computed: {
    OppositeMobileV() {
      return !this.mobileV;
    }
  },
  components: {
    vSearch,
    vHeader
  },
  mounted() {
    this.axios({
      method: "TREE",
      url: "/",
      data: null
    })
      .then(res => {
        this.data = this.changeName(res.data.Nodes);
      })
      .catch(err => {
        console.log("没有文件");
      });
  },
  methods: {
    handleNodeClick(data) {
      console.log(data);
      if (data.className !== "folder") {
        if (this.$store.state.fixed === false)
          this.$emit("handleMobile", false);

        if (data.Nodes.length > 0) {
          for (var i = 0; i < data.Nodes.length; i++) {
            var reg = /[^<>/\\\|:""\*\?]+\.\w+$/;
            data.Nodes[i].Name = data.Nodes[i].Path.match(reg)[0];
          }
        }

        var name = data.Path.substring(data.Path.indexOf("/") + 1);

        if (data.className === "pdficon") {
          window.open(window.location.href + name, data.Path);
        }
        document.title = data.Path;
        this.$router.push({ name: "content", params: { pageId: name } });

        this.$store.commit("changeValue", data);
      }
    },
    handleWidth(e) {
      let self = this;
      let clientX = e.clientX;
      let dw = e.target.offsetWidth;

      // if (
      //   clientX < e.target.offsetWidth - 20 ||
      //   clientX > e.target.offsetWidth + 20
      // ) {
      // } else {
      e.target.onmousemove = function(e) {
        this.W = e.clientX - clientX;
        //self.$store.state.number = dw + this.W;
        e.target.style.width = dw + this.W + "px";
      };
      document.onmouseup = function() {
        //self.$store.state.number = "";
        //document.onmousedown = null;
        e.target.onmousemove = null;
      };
      //}
    },
    handleAside() {
      this.$refs.aside.style.display = "block";
      this.$refs.menu.style.display = "none";
      this.$store.state.splitbar = true;
      this.$emit("handleMobile", true);
    },
    handleHide() {
      if (this.$store.state.fixed === true) {
        return;
      }
      this.$refs.aside.style.display = "none";
      this.$refs.menu.style.display = "block";
      this.$store.state.splitbar = false;
      this.$emit("handleMobile", false);
    },
    handleFixed(value) {
      this.$store.state.fixed = value;
    },
    changeName(data) {
      console.log(data);
      if (data.length == 0) {
        return;
      }
      for (var i = 0; i < data.length; i++) {
        console.log(222);
        if (data[i].Nodes.length == 0) {
          console.log(333);
          var reg = /[^<>/\\\|:""\*\?]+\.\w+$/;
          if (data[i].Path.match(reg)) {
            data[i].Name = data[i].Path.match(reg)[0];
            console.log(555, data[i].Name);
            return data;
          }
        } else {
          continue;
        }
      }
      return this.changeName(data[i].Nodes);
    },
    // el-tree 添加自定义图标
    renderContent: function(h, { node, data, store }) {
      if (data.Path.indexOf(".md") != -1) {
        data.className = "mdicon";
      } else if (data.Path.indexOf(".pdf") != -1) {
        data.className = "pdficon";
      } else {
        data.className = "folder";
      }
      return (
        <span class="icon">
          <i class={data.className} />
          <span style="padding-left: 4px;">{node.label}</span>
        </span>
      );
    }
  }
};
</script>

<style lang="stylus" scoped>
.aside-container >>> .el-tree-node__content {
  border-left: 1px solid #00a0e9;
  padding-left: 10px !important;
  padding-top: 8px;
  padding-bottom: 8px;
}

.aside-container >>> .el-tree-node>.el-tree-node__children {
  padding-left: 10px;
}

.aside-container >>> .el-tree-node__expand-icon {
  position: absolute;
  right: 20px;
}

.aside-container >>> .el-tree-node:focus > .el-tree-node__content {
  border-left: 3px solid #00a0e9;
  color: #00a0e9;
  font-weight: 800;
}

.aside-container {
  display: flex;
  flex-direction: column;
  width: 250px;
  min-width: 200px;
  overflow: hidden;
  height: 100%;
  position: fixed;
}

@media screen and (max-width: 768px) {
  .aside-container {
    display: flex;
    flex-direction: column;
    width: 100%;
    max-width: 600px;
    overflow: hidden;
    height: 100%;
    position: fixed;
  }
}
</style>
<style>
.tree {
  overflow-y: scroll;
  margin-right: -10px;
}
.menu-con {
  background-color: rgb(242, 245, 247);
  box-shadow: rgba(118, 118, 118, 0.11) 2px 0px 5px 0px;
  opacity: 1;
  height: 52px;
  line-height: 2;
  position: absolute;
  left: 0;
  text-align: center;
  top: 33%;
  width: 14px;
  z-index: 2;
  cursor: pointer;
  border-radius: 0px 4px 4px 0px;
  border-width: 1px 1px 1px;
  border-style: solid solid solid none;
  border-color: rgb(224, 228, 231) rgb(224, 228, 231) rgb(224, 228, 231);
  border-image: initial;
  border-left: none;
  padding: 6px;
  transition: right 0.25s ease-in 0.2s, opacity 0.35s ease-in 0.2s;
}

/* .aside-container .el-tree-node__content:hover{
	background:#b1eee9 url(../assets/iconfont/-.png) no-repeat left 4px;
	color:#00a0e9;
	font-weight :800
}  */
.icon .folder {
  width: 36px;
  height: 36px;
  line-height: 36px;
  background: url(../assets/iconfont/文件夹.svg) no-repeat left center;
}
.icon .mdicon {
  width: 36px;
  height: 36px;
  line-height: 36px;
  background: url(../assets/iconfont/markdown-fill.svg) no-repeat left center;
}
.el-tree-node__content .icon {
  height: 36px;
  line-height: 36px;
  display: flex;
}
</style>
