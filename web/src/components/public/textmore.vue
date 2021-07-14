<template>
  <div>
    <el-form :model="ruleForm" :inline="true" ref="ruleForm" label-width="65px" size="medium">
      <el-row
          v-for="(item,index) in ruleForm.formList"
          :key="index"
          style="border-bottom: 1px solid #f0f0f0;padding: 10px;"
      >
        <el-form-item
            label="参数名"
            :prop="'formList.' + index + '.fieldName'"
            :rules="[{ required: true, message: '参数名不能为空',trigger: 'change'},{max: 32, message: '不超过32个字符', trigger: 'change'}]"
        >
          <el-input v-model="item.fieldName" placeholder="请输入参数名" />
        </el-form-item>
        <el-form-item
            label="时间"
            :prop="'formList.' + index + '.fieldTime'"
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
        <el-form-item
            label="描述"
            :prop="'formList.' + index + '.fieldDesc'"
            :rules="[{ required: true, message: '描述不能为空',trigger: 'change'}]"
        >
          <el-select clearable v-model="item.fieldDesc" placeholder="请选择描述">
            <el-option
                v-for="item in options"
                :key="item.value"
                :label="item.label"
                :value="item.value"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item
            label="排序"
            :prop="'formList.' + index + '.fieldSort'"
            :rules="rules.fieldSort"
        >
          <el-input type="number" v-model="item.fieldSort" autocomplete="off" placeholder="请输入排序" />
        </el-form-item>
        <el-button
            type="danger"
            v-if="ruleForm.formList.length > 1"
            size="medium"
            @click="removeRow(index)"
        >删除</el-button>
      </el-row>
      <el-row>
        <el-button type="primary" size="medium" @click="addRow">新增属性</el-button>
        <el-button type="primary" size="medium" @click="submit('ruleForm')">提交</el-button>
      </el-row>
    </el-form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      // select下拉项
      options: [
        {
          value: "选项1",
          label: "黄金糕"
        },
        {
          value: "选项2",
          label: "双皮奶"
        }
      ],
      ruleForm: {
        formList: [
          {
            fieldName: "",
            fieldSort: "",
            fieldDesc: "",
            fieldTime: ""
          }
        ]
      },

      // 新增表单的验证规则
      rules: {
        fieldSort: [
          { required: true, message: "请输入排序", trigger: "change" },
          {
            validator: (rule, value, callback) => {
              if (value < 0) {
                callback(new Error("必须大于0"));
              } else if (value.length > 5) {
                callback(new Error("不超过5位数字"));
              } else if (!value) {
                callback(new Error("排序不能为空"));
              } else {
                callback();
              }
            },
            trigger: "change"
          }
        ]
      }
    };
  },
  methods: {
    addRow() {
      this.ruleForm.formList.push({
        fieldName: "",
        fieldSort: "",
        fieldDesc: "",
        fieldTime: ""
      });
    },
    // 删除属性列
    removeRow(index) {
      this.ruleForm.formList.splice(index, 1);
    },
    // 提交
    submit(formName) {
      this.$refs[formName].validate(valid => {
        console.log(this.ruleForm);
        if (valid) {
          alert("submit!");
        } else {
          return false;
        }
      });
    }
  }
};
</script>
<style></style>