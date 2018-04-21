using System;
using System.Reflection;

namespace typeflection
{
    class BaseModel
    {
        // 3)
        public string GetTypeName()
        {
            return this.GetType().Name;
        }
        // 4)
        public string TypeofName()
        {
            return typeof(BaseModel).Name;
        }
        // 7)
        public static string GenericStaticName<T>(T t)
        {
            return t.GetType().Name;
        }
        // 8)
        public static string StaticName()
        {
            return MethodBase.GetCurrentMethod().DeclaringType.Name;
        }
    }
}
