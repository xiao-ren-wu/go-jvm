public class MyObject{
    public static int staticVar;
    public int instanceVar;

    public static void main(String args[]){
        int x = 32768;                      //ldc
        MyObject myObj = new MyObject();    //new
        MyObject.staticVar = x;             //putstatic
        x = MyObject.staticVar;             //getstatic
        myObj.instanceVar = x;              //putfield
        Object obj = myObj;                 //getfield
        if (obj instanceof MyObject){       //
            myObj = (MyObject) obj;         //instanceof
            System.out.println(myObj.instanceVar);
        }
    }
}
