<template>
  <div>
    <div>
      <div id="title">
        <a-popconfirm
          title="确定不保存就返回吗?"
          ok-text="Yes"
          cancel-text="No"
          @confirm="$router.go(-1)"
        >
          <a-button><a-icon type="left" />返回</a-button>
        </a-popconfirm>
        <a-input placeholder="请输入标题" v-model="title" />
        <a-input-search
          placeholder="请输入标签id"
          @search="save"
        >
          <a-button slot="enterButton"> 发布 </a-button>
        </a-input-search>
      </div>
      <br />
      <mavon-editor
        id="edit"
        ref="edit"
        v-model="value"
        @save="save"
        @imgAdd="imgAdd"
        @imgDel="imgDel"
      />
    </div>
  </div>
</template>

<script>

export default {
  name: "Edit",
  data() {
    return {
      value: null,
      img_file: null,
      title: null,
      select: false,
    };
  },
  methods: {
    // 保存文章
    async save(tid=0) {
      console.log(this.title, tid);
      if(tid<1)  {
        this.$message.info("文章的标签id不饿能小于1哦")
        return
      }
      if (this.title == null || this.title == "") {
        return this.$message.warning("标题不能为空哦");
      }
      if (this.img_file != null) {
        this.$message.info("正在上传图片");
        // await this.uploadimg();
      }
      const { title, value: body } = this;
      try {
        const res = await this.$api.Article.save({
          title,
          content:body,
          tag_id: Number(tid),
        });
        console.log("保存文章的结果", res);
        const { code, msg } = res;
        if (code === 200) {
          this.$message.success("发布成功");
          this.$router.push("/article");
        } else {
          this.$message.warning(msg);
        }
      } catch (error) {
        this.$message.warning("发生错误");
      }
    },
    imgAdd(pos, $file) {
      console.log("添加图片", pos, $file);
      // 缓存图片信息
      this.img_file[pos] = $file;
    },
    imgDel(pos) {
      delete this.img_file[pos[0]];
    },
    async uploadimg() {
      // 第一步.将图片上传到服务器.
      var formdata = new FormData();
      console.log(this.img_file[1]);
      for (let _img in this.img_file) {
        // console.log(_img);
        formdata.append(_img, this.img_file[_img]);
      }

      // console.log("图片数据", formdata)
      const result = await this.$api.Upload.img(formdata);
      // console.log("上传图片的结果", result)
      if (result.status === 201) {
        // 替换源文本的图片链接
        this.$message.success("图片上传成功");
        const $img = this.$refs.edit;

        // 数据清洗
        const imgs = Object.keys(result.data.files).map((key) => {
          // console.log(key)
          return [Number(key), result.data.files[key][0]];
        });
        // console.log(imgs)
        for (let img of imgs) {
          $img.$img2Url(img[0], img[1].path);
        }
      } else {
        this.$message.error("上传失败");
      }
    },
  },
};
</script>

<style lang='less' scoped>
#title {
  margin-top: 10px;
  display: flex;
}
#edit {
  height: 800px;
  z-index: 0;
}
</style>