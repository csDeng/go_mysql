<template>
  <div>
    <a-card title='标签管理' :style="{
      height:'740px'
    }">
    <a slot='extra' id='addbtn' @click='$data.visible = true'>添加标签</a>
    <a-modal
      
      title="添加标签"
      :visible="visible"
      @ok="addCategory"
      @cancel="$data.visible = false"
    >
      <a-space>
        <a-space>
          <a-input  v-model='obj.title' placeholder='请输入标签标题'/>
          <!-- <a-input v-model='obj.description' placeholder='请输入分类描述'/> -->
        </a-space>
      </a-space>
    </a-modal>
    <a-table bordered :data-source="dataSource" :columns="columns" rowKey="name"
        :pagination="{
          position:'bottom',
          pageSize:10,
          total:total,
          onChange:pageChage
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
        <!-- <a-divider type="vertical" /> -->
      </template>
    </a-table>
    </a-card>
  </div>
</template>
<script>
import columns from './columns.js';

export default {
  data() {
    return {
      columns,
      dataSource: [], 
      api: this.$api.Tags,
      visible: false,
      total:0,
      obj: {
        title: null,
        description: null
      }
    };
  },
  mounted(){
    this.get_list()
  },
  methods: {
    get_list(pid=0){
      this.api.get_list(pid).then(r=>{
        let { lists, total} = r.data
        lists = lists.map(item=>{
          if(!item.created_by) {
            item.created_by = "默认"
          }
          if(!item.modified_by) {
            item.modified_by = "暂时无人修改"
          }
          return item
        })
        this.dataSource = lists
        this.total = total

      })
    },
    pageChage(e){
      this.get_list(e)
    },
    addCategory(){
      console.log('add tag=>',this.obj)
      this.api.add({name:this.obj.title}).then(r=>{
        if(r.status === 201 ){
          this.$message.success('添加成功')
          this.get_list()
          this.visible = false
        }else{
          this.$message.error("添加失败")
        }
      })
    },
    onDelete(record) {
      console.log(record)
      this.api.delete(record.id).then(r=>{
        if(r.status === 200 ) {
          this.$message.success('删除成功')
          this.get_list()
        }else{
          this.$message.warning('发生错误')
        }
      })
    }
  },
};
</script>
<style scoped lang='less'>
#addbtn{
  margin-bottom: 20px;
}
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