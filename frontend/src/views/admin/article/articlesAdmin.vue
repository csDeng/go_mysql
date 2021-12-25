<template>
  <a-card
    title="文章管理"
    :style="{
      height: '740px',
    }"
  >
    <a slot="extra" class="editable-add-btn" @click="$router.push('/Edit')"
      >添加文章</a
    >
    <a-table bordered :data-source="dataSource" :columns="columns" rowKey="id" 
        :pagination="{
          position:'bottom',
          pageSize:10,
          total:total,
          onChange:pageChange
        }"
    >
      <template slot="operation" slot-scope="text, record">
        <a-popconfirm
          v-if="dataSource.length"
          title="确认删除吗?"
          @confirm="() => onDelete(record)"
        >
          <a href="javascript:;">删除</a>
        </a-popconfirm>
      </template>
      <a-pagination size="small" :total="50" />
    </a-table>
  </a-card>
</template>
<script>
import columns from "./columns.js";
export default {
  data() {
    return {
      dataSource: [],
      count: 2,
      columns,
      api: this.$api.Article,
      total: 0,
    };
  },
  mounted() {
    this.get();
  },
  methods: {
    get(pid=0) {
      this.api.get_list(pid).then((r) => {
        console.log(r);
        let { lists, total } = r.data;
        this.total = total;
        this.dataSource = lists.map((item) => {
          item.body = item.content.substring(0, 10);
          if (!item.created_by) {
            item.created_by = "默认";
          }
          return item;
        });
        // console.log('获取文章的结果',this.dataSource)
      });
    },
    pageChange(e) {
      this.get(e)
    },
    async onDelete(key) {
      const res = await this.$api.Article.del(key.id);
      console.log(res);
      const { code, msg } = res;
      if (code === 200) {
        this.get();
        await this.$message.success("删除成功");
      } else {
        this.$message.error(msg);
      }
    },
  },
};
</script>
<style scoped lang='less'>
.editable-cell {
  position: relative;
}

.editable-cell-input-wrapper,
.editable-cell-text-wrapper {
  padding-right: 24px;
}

.editable-cell-text-wrapper {
  padding: 5px 24px 5px 5px;
}

.editable-cell-icon,
.editable-cell-icon-check {
  position: absolute;
  right: 0;
  width: 20px;
  cursor: pointer;
}

.editable-cell-icon {
  line-height: 18px;
  display: none;
}

.editable-cell-icon-check {
  line-height: 28px;
}

.editable-cell:hover .editable-cell-icon {
  display: inline-block;
}

.editable-cell-icon:hover,
.editable-cell-icon-check:hover {
  color: #108ee9;
}

.editable-add-btn {
  margin-bottom: 8px;
}
</style>