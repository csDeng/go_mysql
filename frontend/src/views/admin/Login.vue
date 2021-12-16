<template>
  <div>
      <a-card title="登录" id='main' >
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
            <a-input-password
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
            </a-input-password>
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
          </a-form-item>
        </a-form>
      </a-card>
  </div>
</template>

<script>

export default {
  name: "Login",
  beforeCreate() {
    this.form = this.$form.createForm(this, { name: "normal_login" });
  },
  methods: {
    handleSubmit(e) {
      e.preventDefault();
      this.form.validateFields( async (err, values) => {
        if (!err) {
          const { userName: username, password } = values
          let res = await this.$api.User.AdminLogin({
            username,
            password
          })
          const { data, status } = res
          console.log("登录结果res",res)
          if( status === 200 ){
            this.$message.success("登录成功").then(()=>{
              this.$store.dispatch("login", data ).then(() => {
                this.$router.push("/index");
              });
            })
          }else{
            this.$message.error('登录失败')
          }
          }
      });
    },
  },
};
</script>
<style lang='less' scoped>
#main{
  max-width: 600px;
  margin: 0 auto;
  position: relative;
  top: 25vh;
  border-radius: 8px;
}
</style>