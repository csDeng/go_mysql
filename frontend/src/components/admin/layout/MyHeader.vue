<template>
  <div>
    <a-layout-header
      :style="{
        background: '#fff',
        width: '100%',
        zIndex: 100,
        position: 'fixed',
      }"
      id="layoutHeader"
    >
      <div class="userInfoWrap">
        <a-popover>
          <template slot="content">
              <div v-show="$store.state.isLogin">
                  <a-button @click='logout'>退出登录</a-button>
              </div>
              <div v-show="!$store.state.isLogin">
                  <a-button @click='login'>登录</a-button>
              </div>
          </template>
          <a-space>
                <a-avatar icon="user" />
                <span>{{$store.state.username}}</span>
          </a-space>
         
        </a-popover>
      </div>
    </a-layout-header>
  </div>
</template>

<script>
export default {
  name: "MyHeader",
  data() {
    return {};
  },
  methods: {
      async logout() {
            await this.$store.dispatch('logout')
      },
      login() {
          window.location = "/login"
      }
  },
};
</script>


<style lang='less' scoped>
#layoutHeader {
  display: flex;

  .userInfoWrap {
    position: fixed;
    right: 100px;
  }
}
</style>