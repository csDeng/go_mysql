<template>
  <div>
    <div class="selectTags">
      <a-space direction="vertical" align='center' >

        <a-popconfirm title="确认是当前标签了吗？" ok-text="Yes" cancel-text="No" @confirm='confirmSave'>
          <a-button type="primary" ghost >确认发布</a-button>
        </a-popconfirm>
      </a-space>
    </div>
  </div>
</template>

<script>
export default {
  name: "GetExt",
  data() {
    return {
      tags: [],
      categories:[],
      articleCategories:[],
      inputVisible: false,
      inputValue: "",
    };
  },
  async mounted(){
    // this.getCategories()
    this.getTags()
  },
  methods: {
    addtag(name){
      this.$api.Tags.add({name}).then(r=>{
        if(r.status === 201){
          this.$message.success("添加成功")
          this.getTags()
        }else{
          this.$message.error("添加失败")
        }
      }).then(()=>{
        this.inputVisible = false,
        this.inputValue = ""
      })
    },

    getTags(){
      this.$api.Tags.get().then(r=>{
        console.log('get tags=>',r)
        const { data:{ lists }, code} = r
        const tags = lists
        console.log(tags)
        if( code ===200 ){
          this.tags = tags.map(r=>r.name)
          // console.log(this.tags)
        }else{
          this.$message.error( "获取标签失败" )
        }
      },()=>{
        this.$message.errror("网络出错啦")
      })
    },
    handleClose(removedTag) {
      const tags = this.tags.filter((tag) => tag !== removedTag);
      // console.log(tags);
      this.tags = tags;
    },

    showInput() {
      this.inputVisible = true;
      this.$nextTick(function () {
        this.$refs.input.focus();
      });
    },

    handleInputChange(e) {
      this.inputValue = e.target.value;
    },

    handleInputConfirm() {
      const inputValue = this.inputValue;
      this.addtag(inputValue)
    },


    // 确认发布
    confirmSave(){
        this.$emit("comfirm", { tags:this.tags } )
    }
  },
};
</script>

<style lang='less' scoped>
.select {
  min-height: 150px;
  width:210px;
}
.send {
  display: flex;
}
</style>