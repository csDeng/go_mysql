<template>
  <div id="footer">
    <a-button @click="backTop">回到顶部</a-button>
    <p>
      <a-icon type="copyright" />
      2021-2050 flask  <br />
      <!-- <a-icon type="link" />dcs 版权所有 -->
    </p>
    <p>本博客已启动 {{ day }} 天</p>
    <p>粤ICP备2021082063号</p>
  </div>
</template>

<script>

export default {
  name: "Footer",
  data() {
    return {
      day: 0,
    };
  },
  created() {
    this.$api.Utils.getDate().then((r) => {
      this.day = r.data.day;
    });
  },
  mounted(){

  },
  methods: {
    backTop() {
      document.body.scrollTop = document.documentElement.scrollTop = 0;
    }
  }
};
</script>

<style lang='less' scoped>
#footer {
  padding: 20px;
  text-align: center;
}
</style>