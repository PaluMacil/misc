using System;

namespace typeflection
{
    class Model : BaseModel
    {
        // 5)
        internal string MyGetTypeName()
        {
            return this.GetType().Name;
        }
        // 6)
        internal string MyTypeofName()
        {
            return typeof(Model).Name;
        }
    }
}
