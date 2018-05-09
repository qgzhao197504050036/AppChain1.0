/*ԭ����1�������ڽ�������Լʹ��Sale�������ǼǴ��۱Һ�Լ��ַ���������۸񣬹ҵ���2����Ҳ�ѯ��������ֱ�Ӷ�������������������ȷ���ļ۸񣬶��������������ܺ�Լ���ͽ��ף����ܺ�Լ�յ��󣬽���Ӧ�����ʲ�ת������ʻ������ʽ��͵������ʻ������������������������ݽ��м����ɾ����3��������Ҫ����������ֱ�ӽ����׺ż��������������ܺ�Լ�����ܺ�Լ��֤��ȡ����������Ȩ���������������������������*/
pragma solidity ^0.4.18;

interface SHOPRC30{
      function transferFrom(address _from, address _to, uint256 _value) public;
      function approveAndCall(address _spender, uint256 _value, bytes _extraData) public;
} 

contract DeExChange{
    //ExChange things' Information. �����������Ʒ���ݽṹ
    struct ExThingInfo {
          address owner; //the owner's account. �������ʺ�
          SHOPRC30 SaleToken;  //the wishing to exchange thing's  token address  ��������Ʒ���Һ�Լ��ַ
          uint256 amount; //the amount of withing to exchange thiing ��������Ʒ����
          SHOPRC30 ExToken;//=address(0x0000000000000000000000000000000000000000);//the Token wish to exchange ϣ�������ı�
          uint256 price;  //the price of the things per unit ��Ʒ����
          string AttachInfo; //additional information ��ע�򸽼���Ϣ��
        }
 address public Manager; //�������Ա  
  mapping(bytes32=>ExThingInfo)  public Sales; //������������
  SHOPRC30  SaleToken;   //���ö���Ϊ���۴��Һ�Լ
  SHOPRC30  ExToken;   //���ö���Ϊ���������Һ�Լ

 //�������������볷���¼�
 event SaleEvent(bytes32 salehash,address saleToken,address saler, uint256 Amount,address exToken,uint256 price);
 event ApproveEvent(address _from, uint256 _value, address _token, bytes _extraData);
 event BuyEvent(bytes32 salehash,address saleToken,address buyer, uint256 Amount,address exToken,uint256 price);
 event WithdrawEvent(bytes32 salehash);

//�������κ���������Լ����Ա���޸�
        modifier onlyManager {
         require(msg.sender ==Manager);
         _;
     }
//���캯�����������Ա
function DeExChange(
         ) public {
        Manager=msg.sender;    
}

/* Saler send the sale information and transfer the owner power to contract ���ҹҵ������ʲ������Խ�������Լ��Ȩ*/
  function Sale(address _SaleToken,uint256 Amount,address _ExToken,uint256 Price,string Attachstr)  external returns (bool success){
        // ��������Լ��Ȩ�����������еǼǡ�
             SaleToken=SHOPRC30(_SaleToken);   //ʵ�������Ҷ���
             ExToken=SHOPRC30(_ExToken);   //ʵ�������Ҷ���
             SaleToken.approveAndCall(this,Amount,bytes(Attachstr));  //�Խ�������Լ��Ȩ
             bytes32 Saleshash=sha256(_SaleToken,tx.origin,Amount,_ExToken,Price,Attachstr);  //��������Hash������
             Sales[Saleshash]=ExThingInfo(tx.origin,SaleToken,Amount,ExToken,Price, Attachstr);   //������������
             SaleEvent(Saleshash,_SaleToken,tx.origin,Amount,_ExToken,Price);   //���������¼�
             return true;
     }

/* Buyer send money to the contract and get the Token. ��ҷ����ʽ�����ܺ�Լ���������Ӧ�����ʲ�����*/
  function Buy(bytes32 SalesIndex,uint256 Amount) payable public {
        // ���Լ�����ʽ𣬲���ô�������Ȩ�� 
           require(Amount<=Sales[SalesIndex].amount);    //���������Ƿ�С�ڻ������������
           ExToken=Sales[SalesIndex].ExToken;
           if (address(ExToken)==0x0000000000000000000000000000000000000000){
                   require(msg.value==Sales[SalesIndex].price*Amount);  //Ҫ���͵ײ�ҵ��������������Լ۸�
                   Sales[SalesIndex].owner.transfer(Sales[SalesIndex].price*Amount*999/1000);  //�����ҷ������յ�����
                   Manager.transfer(Sales[SalesIndex].price*Amount*1/1000);  //�����������ʺ�֧�����׷ѡ�
            }else{
           //        ExToken=SHOPRC30(Sales[SalesIndex].ExToken;
                   ExToken.approveAndCall(this,Sales[SalesIndex].price*Amount,bytes(Sales[SalesIndex].AttachInfo));  //�����ڹ�����۱ҵĴ��ҶԽ�������Լ��Ȩ��
                   ExToken.transferFrom(msg.sender,Sales[SalesIndex].owner,Sales[SalesIndex].price*Amount*999/1000);//������ת�ƴ�������������Ȩ��
                   ExToken.transferFrom(msg.sender,Manager,Sales[SalesIndex].price*Amount*1/1000);//�����������ʺ�֧�����׷�;
            }   
           SaleToken=Sales[SalesIndex].SaleToken;  //������Һ�ԼΪ��ǰ���۴��ҡ�
           SaleToken.transferFrom(Sales[SalesIndex].owner,msg.sender,Amount);  //�����ת�ƴ��۴�������Ȩ
           Sales[SalesIndex].amount-=Amount;   //�������м�ȥ��������
           if (Sales[SalesIndex].amount==0){                     //�����������Ϊ0������������������ɾ��
                 delete Sales[SalesIndex];
           } 
           BuyEvent(SalesIndex,SaleToken,msg.sender,Amount,ExToken,Sales[SalesIndex].price);  //�������¼�
      }
function Withdraw(bytes32 SalesIndex) public {
        // ָ�������ţ������ҳ��ء� 
           bytes memory b = new bytes(0);
           require(Sales[SalesIndex].owner==msg.sender);   //ֻ�����������߲��ܳ�����
           SaleToken=Sales[SalesIndex].SaleToken;  //������Һ�ԼΪ��ǰ���۴��ҡ�
           SaleToken.approveAndCall(this,0,b);  //ȡ����������Ȩ     
           delete Sales[SalesIndex]; //������������ɾ��������.
           WithdrawEvent(SalesIndex);   //���������¼�
      } 
 function receiveApproval(address _from, uint256 _value, address _token, bytes _extraData) public{
      ApproveEvent(_from,_value,_token,_extraData);  //�ɴ�����Ȩ�ɹ�ʱ���ã�������Ȩ�ɹ��¼�
  }
//Management power thansfer. ����Ȩת�ƣ������ɵ�ǰ��Լ����Ա����
  function transferManagement(address newManager) onlyManager public {
               Manager=newManager;
       }
}