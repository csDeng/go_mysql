<template>
  <div>
    <a-layout-sider
      v-model="collapsed"
      :theme="theme"
      collapsible
      breakpoint="lg"
      :style="{
        height: '100vh',
        width:'200px',
        overflowY: 'auto',
      }"
    >
      <div id="logo">
        <img src="../../../img/logo.jpg" @click="goto('/')" :title="title" />
      </div>
      <a-menu :theme="theme" mode="inline" :default-selected-keys="['1']">
        <a-menu-item key="1" @click="goto('/')">
          <a-icon type="dashboard" />
          <span>仪表盘</span>
        </a-menu-item>
        <a-menu-item key="2" @click="goto('/user')">
          <a-icon type="user" />
          <span>用户管理</span>
        </a-menu-item>
        <a-menu-item key="3" @click="goto('/article')">
          <a-icon type="book" />
          <span>文章管理</span>
        </a-menu-item>
        <a-menu-item key="4" @click="goto('/category')">
          <a-icon type="database" />
          <span>标签管理</span>
        </a-menu-item>
      </a-menu>
    </a-layout-sider>
  </div>
</template>

<script>
export default {
  name: "Slider",
  data() {
    return {
      collapsed: false,
      theme: "light",
      title: "博客后台管理",
      logoSrc: "../../../img/logo.jpg",
    };
  },
  methods: {
    goto(param) {
      const temp = this.$router.history.current.fullPath;
      if (temp === param) return;
      this.$router.push(param);
    },
  },
};
</script>



<style lang='less'>
#logo {
  img {
    width: 50px;
    height: 50px;
    border-radius: 50%;
    margin-left: 30%;

    &:hover {
      -webkit-animation: logo 2s linear infinite;
      animation: logo 2s linear infinite;
    }
  }
}
@keyframes logo {
  0% {
    transform: rotatey(0deg);
  }
  100% {
    transform: rotatey(360deg);
  }
}

#components-layout-demo-custom-trigger {
  min-height: 75vh;
}
#components-layout-demo-custom-trigger .trigger {
  font-size: 18px;
  line-height: 64px;
  padding: 0 24px;
  cursor: pointer;
  transition: color 0.3s;
}

#components-layout-demo-custom-trigger .trigger:hover {
  color: #1890ff;
}

#components-layout-demo-custom-trigger .logo {
  height: 32px;
  background: rgba(255, 255, 255, 0.2);
  margin: 16px;
}
</style>