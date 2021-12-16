export default [
    {
      title:'id',
      dataIndex:'id',
      key:'id'
    },
    {
      title: '用户名',
      dataIndex: 'username',
      key:'username',
      scopedSlots: { customRender: 'username' },
    },
    {
      title: '权限',
      dataIndex: 'level',
      key: 'level',
    },
    {
      title: '操作',
      dataIndex: 'operation',
      scopedSlots: { customRender: 'operation' },
    },
];