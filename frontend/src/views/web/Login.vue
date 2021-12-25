<template>
  <div>
    <div id="app">
      <div id="main1">
        <a-card >
          <a-form
            id="login"
            :form="form"
            class="login-form"
            @submit="handleSubmit"
          >
            <a-form-item label="用户名" labelAlign="left" required hasFeedback>
              <a-input
                v-decorator="[
                  'userName',
                  {
                    rules: [
                      {
                        required: true,
                        message: 'Please input your username!',
                      },
                    ],
                  },
                ]"
                placeholder="Username"
              >
                <a-icon
                  slot="prefix"
                  type="user"
                  style="color: rgba(0, 0, 0, 0.25)"
                />
              </a-input>
            </a-form-item>
            <a-form-item label="密码" labelAlign="left" required hasFeedback>
              <a-input
                v-decorator="[
                  'password',
                  {
                    rules: [
                      {
                        required: true,
                        message: 'Please input your Password!',
                      },
                    ],
                  },
                ]"
                type="password"
                placeholder="Password"
              >
                <a-icon
                  slot="prefix"
                  type="lock"
                  style="color: rgba(0, 0, 0, 0.25)"
                />
              </a-input>
            </a-form-item>
            <a-form-item>
              <a-checkbox
                v-decorator="[
                  'remember',
                  {
                    valuePropName: 'checked',
                    initialValue: true,
                  },
                ]"
              >
                Remember me
              </a-checkbox>
              <br />
              <a-button type="primary" html-type="submit"> 登录 </a-button>
              Or
              <a href="javascript:;" @click="$router.push('/Register')">
                注册
              </a>
            </a-form-item>
          </a-form>
        </a-card>
      </div>
    </div>
     <vue-particles
     id="particles-js"
          color="#409EFF"
          :particleOpacity="0.7"
          :particlesNumber="60"
          shapeType="circle"
          :particleSize="6"
          linesColor="#409EFF"
          :linesWidth="1"
          :lineLinked="true"
          :lineOpacity="0.4"
          :linesDistance="150"
          :moveSpeed="5"
          :hoverEffect="true"
          hoverMode="grab"
          :clickEffect="true"
          clickMode="push">
</vue-particles>

  </div>
</template>

<script>
import VueParticles from 'vue-particles'  
import Vue from 'vue'
Vue.use(VueParticles)

export default {
  name: "Login",
  beforeCreate() {
    this.form = this.$form.createForm(this, { name: "normal_login" });
  },
  methods: {
    handleSubmit(e) {
      e.preventDefault();
      this.form.validateFields(async (err, values) => {
        if (!err) {
          const { userName: username, password } = values;
          try {
            let res = await this.$api.User.login({
              username,
              password,
            });
            const { code: status, data } = res;
            // console.log("登录结果res", res);
            if (status === 200) {
              this.$message.success("登录成功").then(() => {
                this.$store.dispatch("login", data).then(() => {
                  this.$router.push("/Index");
                });
              });
            } else {
              this.$message.error("登录失败");
            }
          } catch (e) {
            console.log("login fail", e);
          }
        }
      });
    },
  },
};
</script>
<style lang='less' scoped>

#main1 {
  position:fixed;
  top:150px;
  left:50%;
  margin-left:-300px;
  // margin: 150px auto;
  border-radius: 8px;
  width: 600px;
  z-index:999;
}
#particles-js {
  height: calc(100% - 100px);
  position: absolute;
}
</style>