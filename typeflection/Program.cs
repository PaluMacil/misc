using System;

namespace typeflection
{
    class Program
    {
        static void Main(string[] args)
        {
            /* 
                GetType() will always return the actual instantiated type. See example 3 where the method on the base type
                returns the name of the Model object.
            */
            Console.WriteLine("1) typeof(Model): {0}", typeof(Model).Name); // Model
            Console.WriteLine("2) model.GetType(): {0}", new Model().GetType().Name); // Model
            Console.WriteLine("3) inherited from base, model.GetTypeName(): {0}", new Model().GetTypeName()); // Model
            Console.WriteLine("4) inherited from base, model.TypeofName(): {0}", new Model().TypeofName()); // BaseModel
            Console.WriteLine("5) implemented in model, model.MyGetTypeName(): {0}", new Model().MyGetTypeName()); // Model
            Console.WriteLine("6) implemented in model, model.MyTypeofName(): {0}", new Model().MyTypeofName()); // Model
            Console.WriteLine("7) static method in base, on generic BaseModel.GenericStaticName(new Model()): {0}", BaseModel.GenericStaticName(new Model())); // Model
            Console.WriteLine("8) static method in base, called with BaseModel.StaticName(): {0}", BaseModel.StaticName()); // BaseModel
            Console.WriteLine("9) static method in base, called with Model.StaticName(): {0}", Model.StaticName()); // BaseModel
        }
    }
}
