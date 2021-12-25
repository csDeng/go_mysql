<template>
  <div id="page">
    <a-card :bordered="false" title="用户管理" id="user">
      <div id="add" slot="extra">
        <a-space>
          <a-input addonBefore="用户名:" v-model="obj.username" />
          <a-input-password addonBefore="密码:" v-model="obj.password" />
          <a-button @click="addUser">添加用户</a-button>
        </a-space>
      </div>
      <a-table
        :columns="columns"
        :data-source="data"
        bordered
        rowKey="username"
        :pagination="{
          position:'bottom',
          pageSize:10,
          total:total,
          onChange:pageChage
        }"
      >
        <template slot="level" slot-scope="text">
          {{ level[text] }}
        </template>
        <template slot="operation" slot-scope="text, record">
          <div class="editable-row-operations">
            <span ref="myedit">
              <a @click="() => edit(record)">编辑</a>
            </span>

            <a-divider type="vertical" />
            <span>
              <a-popconfirm
                title="确认删除该用户吗?"
                ok-text="Yes"
                cancel-text="No"
                @confirm="del(record)"
              >
                <a href="javascript:;">删除</a>
              </a-popconfirm>
            </span>
          </div>
        </template>
      </a-table>
      <a-modal
        :title="'当前编辑用户:' + (cache ? cache.username : '编辑')"
        :visible="visible"
        :confirm-loading="confirmLoading"
        @ok="save"
        @cancel="$data.visible = false"
      >
        <a-space>
                <a-input addonBefore="用户名:" v-model="cache.username" />
                <a-input addonBefore="密码:" v-model="cache.password" />
                <a-select
                  :default-value="cache ? cache.level : '默认权限'"
                  style="width: 120px"
                  @change="select"
                >
                  <a-select-option value="0"> 普通用户 </a-select-option>
                  <a-select-option value="1"> 管理员 </a-select-option>
                </a-select>
              </a-space>
      </a-modal>
    </a-card>
  </div>
</template>

<script>
import columns from "./userColumn.js";
export default {
  name: "userAdmin",
  data() {
    return {
      api: null,
      columns,
      data: [],
      cache: {},
      visible: false,
      confirmLoading:false,
      total:20,
      obj: {
        username: null,
        password: null,
      },
    };
  },
  mounted() {
    this.api = this.$api.User;
    this.get_list();
    window.onresize = () => {
      if (!this.$refs.edit) return;
      if (this.visible) this.$refs.myedit.click();
    };
  },
  methods: {
    get_list(pid=0) {
      this.api.get_list(pid).then((r) => {
        let {lists, total} = r.data
        this.data = lists.map(item=>{
          if(item.level === "1") {
            item.level = "管理员"
          } else {
            item.level = "普通用户"
          }
          return item
        });
        this.total = total;
      });
    },
    pageChage(e){
      this.get_list(e)
    },
    edit(record) {
      // 深拷贝进行数据隔离
      this.cache = { ...record };
      this.cache.level = this.cache.level??0
      this.visible = true;
    },
    select(value) {
      this.cache.level = value;
    },
    save() {
      this.confirmLoading = true
      console.log("edit save=>",this.cache)
      this.api.update(this.cache).then(
        (r) => {
          console.log("添加用户=>", r);
          const {code, msg} = r
          if (code === 200) {
            this.get_list();
            this.$message.success("更新成功");
            this.visible = false;
          }else {
              this.$message.error(msg);
          }
        },
        () => {
          this.$message.error("更新失败");
        }
      );
      this.confirmLoading = false
    },
    cancel() {
      console.log("cancel");
      this.cache = {}
      this.visible = false
    },
    del(val) {
      this.api.del(val).then(
        (r) => {
          if (r.code === 200) {
            this.$message.success("删除成功");
            this.get_list();
          } else {
            this.$message.error(r.msg)
          }
        },
        (e) => {
          console.log(e);
        }
      );
    },
    addUser() {
      const { obj, api } = this;
      let { username, password } = obj;
      if (!username) return this.$message.warning("请填写用户名");
      if (!password) return this.$message.warning("请填写密码");

      api.register(obj).then(
        (r) => {
          console.log("tianjiayonghu =>", r);
          const { code, msg } = r;
          if (code === 200) {
            this.get_list();
            return this.$message.success("添加成功");
          }
          this.$message.warning(msg);
        },
        () => {
          this.$message.error("服务器发生错误");
        }
      );
    },
  },
};
</script>

<style lang='less' scoped>
</style>