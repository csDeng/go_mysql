<template>
  <a-card title='文章管理'>
      <a slot='extra' class="editable-add-btn" @click="$router.push('/Edit')">添加文章</a>
    <a-table bordered :data-source="dataSource" :columns="columns" rowKey='id'>
      <template slot="operation" slot-scope="text, record">
        <a-popconfirm
          v-if="dataSource.length"
          title="Sure to delete?"
          @confirm="() => onDelete(record)"
        >
          <a href="javascript:;">删除</a>
        </a-popconfirm>
      </template>
    </a-table>
  </a-card>
</template>
<script>
import columns from './columns.js';
export default {
  data() {
    return {
      dataSource: [],
      count: 2,
      columns,
      api: this.$api.Article
    };
  },
  mounted(){
    this.get()
  },
  methods: {
    get(){
      this.api.get_list().then(r=>{
        this.dataSource = r.data.map(item=>{
          item.body = item.body.substring(0,10)
          return item
        })
        console.log('获取文章的结果',this.dataSource)
      })
    },

    async onDelete(key) {
      const res = await this.$api.Article.del(key.id)
      console.log(res)
      if(res.status === 200){
        this.get()
        await this.$message.success('删除成功')
      }
      else{
        this.$message.error('删除失败')
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