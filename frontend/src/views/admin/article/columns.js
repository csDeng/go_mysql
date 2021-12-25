export default [
    {
      title: "id",
      dataIndex: "id"
    },
    {
      title: "作者",
      dataIndex: "created_by",
    },
    {
      title: "文章体",
      dataIndex: "body",
    },
    {
      title:"文章所属标签",
      dataIndex:"tag.name"
    },
    {
      title:"软删除",
      dataIndex:"deleted_on"
    },
    {
      title: "操作",
      dataIndex: "operation",
      scopedSlots: { customRender: "operation" },
    },
];