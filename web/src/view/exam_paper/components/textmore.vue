<template>
  <div>
    <el-form :model="ruleForm" :inline="true" ref="ruleForm" label-width="65px" size="medium">
      <el-row
          v-for="(item,index) in ruleForm.list"
          :key="index"
          style="border-bottom: 1px solid #f0f0f0;padding: 10px;"
      >
        <el-form-item
            label="内容"
            :prop="'list.' + index + '.fieldName'"
            :rules="[{ required: true, message: '内容不能为空',trigger: 'change'},{max: 32, message: '不超过32个字符', trigger: 'change'}]"
        >
          <el-input v-model="item.fieldName" placeholder="请输入内容" />
        </el-form-item>
        <el-form-item
            label="时间"
            :prop="'list.' + index + '.fieldTime'"
            :rules="[{ required: true, message: '时间不能为空',trigger: 'change'}]"
        >
          <el-date-picker
              v-model="item.fieldTime"
              type="daterange"
              start-placeholder="开始日期"
              end-placeholder="结束日期"
              value-format="yyyy-MM-dd"
              :default-time="['00:00:00', '23:59:59']"
          ></el-date-picker>
        </el-form-item>

        <el-button
            type="danger"
            v-if="ruleForm.list.length > 1"
            size="medium"
            @click="removeRow(index)"
        >删除</el-button>
      </el-row>
      <el-row>
        <el-button type="primary" size="medium" @click="addRow">新增属性</el-button>
      </el-row>
    </el-form>
  </div>
</template>

<script>
import globalBus from '@/utils/global-bus'
export default {
  data() {
    return {
      ruleForm: {
        list: [
          {
            fieldName: "",
            fieldTime: ""
          }
        ]
      },
    };
  },
  mounted () {
    globalBus.$on('questiontext', this.resString)
  },
  methods: {
    addRow() {
      this.ruleForm.list.push({
        fieldName: "",
        fieldTime: ""
      });
    },
    // 删除属性列
    removeRow(index) {
      this.ruleForm.list.splice(index, 1);
    },

    created(){
      // 注册自定义事件 globalBus.$on 获取数据
      // 发布 globalBus.$emit 发送数据
    },
    submitinfo(formName) {
      this.$refs[formName].validate(valid => {
        globalBus.$emit("extend_textmore",this.ruleForm)
        console.log(this.ruleForm);
        if (valid) {
         // alert("submit!");
        } else {
          return false;
        }
      });
    },
    resString () {
      this.submitinfo('ruleForm')
    },
  }
};
</script>
<style></style>