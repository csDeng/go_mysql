export default [
  {
    title:'id',
    dataIndex:'id',
  },
    {
      title: '标签名称',
      dataIndex: 'name',
      key:'name',
      scopedSlots: { customRender: 'title' },
    },  
    {
    title:'创建者',
    dataIndex:'created_by'
  },{
    title:"上次修改",
    dataIndex:"modified_by"
  },
  {
    title:"是否软删除",
    dataIndex:"deleted_on"
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