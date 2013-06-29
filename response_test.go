package sqs_test

var TestCreateQueueXmlOK = `
<CreateQueueResponse>
  <CreateQueueResult>
    <QueueUrl>http://sqs.us-east-1.amazonaws.com/123456789012/testQueue</QueueUrl>
  </CreateQueueResult>
  <ResponseMetadata>
    <RequestId>7a62c49f-347e-4fc4-9331-6e8e7a96aa73</RequestId>
  </ResponseMetadata>
</CreateQueueResponse>
`

var TestListQueuesXmlOK = `
<ListQueuesResponse>
  <ListQueuesResult>
    <QueueUrl>http://sqs.us-east-1.amazonaws.com/123456789012/testQueue</QueueUrl>
  </ListQueuesResult>
  <ResponseMetadata>
    <RequestId>725275ae-0b9b-4762-b238-436d7c65a1ac</RequestId>
  </ResponseMetadata>
</ListQueuesResponse>
`

var TestGetQueueUrlXmlOK = `
<GetQueueUrlResponse>
 <GetQueueUrlResult>
   <QueueUrl>http://sqs.us-east-1.amazonaws.com/123456789012/testQueue</QueueUrl>
 </GetQueueUrlResult>
 <ResponseMetadata>
   <RequestId>470a6f13-2ed9-4181-ad8a-2fdea142988e</RequestId>
 </ResponseMetadata>
</GetQueueUrlResponse>
`

var TestChangeMessageVisibilityXmlOK = `
<ChangeMessageVisibilityResponse>
 <ResponseMetadata>
  <RequestId>6a7a282a-d013-4a59-aba9-335b0fa48bed</RequestId>
 </ResponseMetadata>
</ChangeMessageVisibilityResponse>
`

var TestChangeMessaveVisibilityBatchXmlOK = `
<ChangeMessageVisibilityBatchResponse>
 <ChangeMessageVisibilityBatchResult>
  <ChangeMessageVisibilityBatchResultEntry>
   <Id>change_visibility_msg_2</Id>
  </ChangeMessageVisibilityBatchResultEntry>
 <ChangeMessageVisibilityBatchResultEntry>
   <Id>change_visibility_msg_3</Id>
 </ChangeMessageVisibilityBatchResultEntry>
 </ChangeMessageVisibilityBatchResult>
 <ResponseMetadata>
  <RequestId>ca9668f7-ab1b-4f7a-8859-f15747ab17a7</RequestId>
 </ResponseMetadata>
</ChangeMessageVisibilityBatchResponse>
`

var TestReceiveMessageXmlOK = `
<ReceiveMessageResponse>
  <ReceiveMessageResult>
    <Message>
      <MessageId>5fea7756-0ea4-451a-a703-a558b933e274</MessageId>
      <ReceiptHandle>MbZj6wDWli+JvwwJaBV+3dcjk2YW2vA3+STFFljTM8tJJg6HRG6PYSasuWXPJB+CwLj1FjgXUv1uSj1gUPAWV66FU/WeR4mq2OKpEGYWbnLmpRCJVAyeMjeU5ZBdtcQ+QEauMZc8ZRv37sIW2iJKq3M9MFx1YvV11A2x/KSbkJ0=</ReceiptHandle>
      <MD5OfBody>fafb00f5732ab283681e124bf8747ed1</MD5OfBody>
      <Body>This is a test message</Body>
      <Attribute>
        <Name>SenderId</Name>
        <Value>195004372649</Value>
      </Attribute>                                                                                                                   
      <Attribute>
        <Name>SentTimestamp</Name>
        <Value>1238099229000</Value>
      </Attribute>
      <Attribute>
        <Name>ApproximateReceiveCount</Name>
        <Value>5</Value>
      </Attribute>
      <Attribute>
        <Name>ApproximateFirstReceiveTimestamp</Name>
        <Value>1250700979248</Value>
      </Attribute>
    </Message>
  </ReceiveMessageResult>
<ResponseMetadata>
  <RequestId>b6633655-283d-45b4-aee4-4e84e0ae6afa</RequestId>
</ResponseMetadata>
</ReceiveMessageResponse>
`

var TestDeleteMessageXmlOK = `
<DeleteQueueResponse>
  <ResponseMetadata>
    <RequestId>b5293cb5-d306-4a17-9048-b263635abe42</RequestId>
  </ResponseMetadata>
</DeleteQueueResponse>
`

var TestSendMessageXmlOK = `
<SendMessageResponse>
  <SendMessageResult>
    <MD5OfMessageBody>fafb00f5732ab283681e124bf8747ed1</MD5OfMessageBody>
    <MessageId>5fea7756-0ea4-451a-a703-a558b933e274</MessageId>
  </SendMessageResult>
  <ResponseMetadata>
    <RequestId>27daac76-34dd-47df-bd01-1f6e873584a0</RequestId>
  </ResponseMetadata>
</SendMessageResponse>
`

var TestSendMessageBatchXmlOK = `
<SendMessageBatchResponse>
<SendMessageBatchResult>
    <SendMessageBatchResultEntry>
        <Id>test_msg_001</Id>
        <MessageId>0a5231c7-8bff-4955-be2e-8dc7c50a25fa</MessageId>
        <MD5OfMessageBody>0e024d309850c78cba5eabbeff7cae71</MD5OfMessageBody>
    </SendMessageBatchResultEntry>
    <SendMessageBatchResultEntry>
        <Id>test_msg_002</Id>
        <MessageId>15ee1ed3-87e7-40c1-bdaa-2e49968ea7e9</MessageId>
        <MD5OfMessageBody>7fb8146a82f95e0af155278f406862c2</MD5OfMessageBody>
    </SendMessageBatchResultEntry>
</SendMessageBatchResult>
<ResponseMetadata>
    <RequestId>ca1ad5d0-8271-408b-8d0f-1351bf547e74</RequestId>
</ResponseMetadata>
</SendMessageBatchResponse>
`

var TestGetQueueAttributesAllXmlOK = `
<GetQueueAttributesResponse>
  <GetQueueAttributesResult>
    <Attribute>
      <Name>VisibilityTimeout</Name>
      <Value>30</Value>
    </Attribute>
    <Attribute>
      <Name>ApproximateNumberOfMessages</Name>
      <Value>0</Value>
    </Attribute>
    <Attribute>
      <Name>ApproximateNumberOfMessagesNotVisible</Name>
      <Value>0</Value>
    </Attribute>
    <Attribute>
      <Name>CreatedTimestamp</Name>
      <Value>1286771522</Value>
    </Attribute>
    <Attribute>
      <Name>LastModifiedTimestamp</Name>
      <Value>1286771522</Value>
    </Attribute>
    <Attribute>
      <Name>QueueArn</Name>
      <Value>arn:aws:sqs:us-east-1:123456789012:qfoo</Value>
    </Attribute>
    <Attribute>
      <Name>MaximumMessageSize</Name>
      <Value>8192</Value>
    </Attribute>
    <Attribute>
      <Name>MessageRetentionPeriod</Name>
      <Value>345600</Value>
    </Attribute>
  </GetQueueAttributesResult>
  <ResponseMetadata>
    <RequestId>1ea71be5-b5a2-4f9d-b85a-945d8d08cd0b</RequestId>
  </ResponseMetadata>
</GetQueueAttributesResponse>
`

var TestGetQueueAttributesSelectiveXmlOK = `
<GetQueueAttributesResponse>
  <GetQueueAttributesResult>
    <Attribute>
      <Name>VisibilityTimeout</Name>
      <Value>30</Value>
    </Attribute>
    <Attribute>
      <Name>DelaySeconds</Name>
      <Value>0</Value>
    </Attribute>
 </GetQueueAttributesResult>
  <ResponseMetadata>
    <RequestId>1ea71be5-b5a2-4f9d-b85a-945d8d08cd0b</RequestId>
  </ResponseMetadata>
</GetQueueAttributesResponse>
`

var TestDeleteMessageBatchXmlOK = `
<DeleteMessageBatchResponse>
    <DeleteMessageBatchResult>
        <DeleteMessageBatchResultEntry>
            <Id>msg1</Id>
        </DeleteMessageBatchResultEntry>
        <DeleteMessageBatchResultEntry>
            <Id>msg2</Id>
        </DeleteMessageBatchResultEntry>
    </DeleteMessageBatchResult>
    <ResponseMetadata>
        <RequestId>d6f86b7a-74d1-4439-b43f-196a1e29cd85</RequestId>
    </ResponseMetadata>
</DeleteMessageBatchResponse>
`

var TestAddPermissionXmlOK = `
<AddPermissionResponse>
    <ResponseMetadata>
        <RequestId>9a285199-c8d6-47c2-bdb2-314cb47d599d</RequestId>
    </ResponseMetadata>
</AddPermissionResponse>
`

var TestRemovePermissionXmlOK = `
<RemovePermissionResponse>
    <ResponseMetadata>
        <RequestId>f8bdb362-6616-42c0-977a-ce9a8bcce3bb</RequestId>
    </ResponseMetadata>
</RemovePermissionResponse>
`

var TestDeleteQueueXmlOK = `
<DeleteQueueResponse>
    <ResponseMetadata>
        <RequestId>6fde8d1e-52cd-4581-8cd9-c512f4c64223</RequestId>
    </ResponseMetadata>
</DeleteQueueResponse>
`

var TestSetQueueAttributesXmlOK = `
<SetQueueAttributesResponse>
    <ResponseMetadata>
        <RequestId>e5cca473-4fc0-4198-a451-8abb94d02c75</RequestId>
    </ResponseMetadata>
</SetQueueAttributesResponse>
`
