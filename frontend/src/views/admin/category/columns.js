export default [
  {
    title:'id',
    dataIndex:'id',
  },
  {
    title:'uid',
    dataIndex:'uid'
  },
    {
      title: '标签名称',
      dataIndex: 'name',
      key:'name',
      scopedSlots: { customRender: 'title' },
    },
    // {
    //   title: '分类描述',
    //   dataIndex: 'description',
    // },
    {
      title: '操作',
      dataIndex: 'operation',
      scopedSlots: { customRender: 'operation' },
    },
];